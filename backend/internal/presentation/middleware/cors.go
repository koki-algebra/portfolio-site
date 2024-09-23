package middleware

import (
	"net/http"
	"strings"

	"github.com/rs/cors"

	"backend/internal/config"
)

func WithCORS(next http.Handler) http.Handler {
	m := cors.New(cors.Options{
		AllowedOrigins: strings.Split(config.Env.ServerAllowOrigins, ","),
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	return m.Handler(next)
}
