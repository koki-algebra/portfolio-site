package middleware

import (
	"log/slog"
	"net/http"
	"time"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
)

func NewCommonMiddlewares(mux http.Handler) http.Handler {
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

	return With(
		mux,
		httplog.RequestLogger(logger),
		chiMiddleware.Heartbeat("/ping"),
		WithCORS,
	)
}
