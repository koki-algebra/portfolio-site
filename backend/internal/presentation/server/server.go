package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sourcegraph/conc/pool"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"backend/internal/config"
	"backend/internal/infrastructure/database"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	sqlDB, err := database.Open(ctx)
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	router, err := newRouter(ctx, sqlDB)
	if err != nil {
		return err
	}

	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", config.Env.ServerPort),
		WriteTimeout:      time.Second * 60,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
		IdleTimeout:       time.Second * 120,
		Handler:           h2c.NewHandler(router, &http2.Server{}),
	}

	pool := pool.New().WithErrors().WithContext(ctx)
	pool.Go(func(ctx context.Context) error {
		slog.InfoContext(ctx, fmt.Sprintf("start HTTP server port: %d", config.Env.ServerPort))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	<-ctx.Done()
	slog.InfoContext(ctx, "stopping API server...")
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	return pool.Wait()
}
