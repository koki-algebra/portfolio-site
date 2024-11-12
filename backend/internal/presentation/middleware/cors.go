package middleware

import (
	"net/http"
	"strings"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"

	"backend/internal/config"
)

func WithCORS(next http.Handler) http.Handler {
	allowedHeaders := connectcors.AllowedHeaders()
	allowedHeaders = append(allowedHeaders, "Authorization")

	m := cors.New(cors.Options{
		AllowedOrigins: strings.Split(config.Env.ServerAllowOrigins, ","),
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: allowedHeaders,
		ExposedHeaders: connectcors.ExposedHeaders(),
		MaxAge:         7200,
	})

	return m.Handler(next)
}
