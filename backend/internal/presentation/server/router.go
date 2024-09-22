package server

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"strings"
	"time"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/rs/cors"

	"backend/internal/config"
	"backend/internal/presentation/middleware"
	"backend/pkg/di"
	"backend/pkg/grpc/gen/user/v1/userv1connect"
)

func newRouter(ctx context.Context, sqlDB *sql.DB) (http.Handler, error) {
	logger := httplog.NewLogger("api-server", httplog.Options{
		LogLevel:         slog.LevelInfo,
		LevelFieldName:   "severity",
		MessageFieldName: "message",
		JSON:             true,
		Concise:          false,
		RequestHeaders:   true,
		TimeFieldFormat:  time.RFC3339,
		TimeFieldName:    "time",
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
	})

	slog.SetDefault(logger.Logger)

	handlers, err := di.InitConnectService(ctx, sqlDB)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	mux.Handle(userv1connect.NewUserServiceHandler(handlers.UserServiceHandler))

	return middleware.With(
		mux,
		httplog.RequestLogger(logger),
		chiMiddleware.Heartbeat("/ping"),
		cors.New(cors.Options{
			AllowedOrigins: strings.Split(config.Env.ServerAllowOrigins, ","),
			AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
			AllowedHeaders: []string{"Authorization", "Content-Type"},
		}).Handler,
	), nil
}
