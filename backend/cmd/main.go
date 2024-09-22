package main

import (
	"context"
	"log/slog"
	"os"

	"backend/internal/config"
	"backend/internal/presentation/server"
)

func main() {
	if err := run(context.Background()); err != nil {
		slog.Error("failed to terminated server", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	if err := config.Init(ctx); err != nil {
		return err
	}

	srv := server.NewServer()

	return srv.Run(ctx)
}
