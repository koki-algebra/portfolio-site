package server

import (
	"context"
	"database/sql"
	"net/http"

	"connectrpc.com/connect"

	"backend/internal/presentation/interseptor"
	"backend/internal/presentation/middleware"
	"backend/pkg/di"
	"backend/pkg/grpc/gen/user/v1/userv1connect"
)

func newRouter(
	ctx context.Context,
	sqlDB *sql.DB,
) (http.Handler, error) {
	handlers, err := di.InitConnectService(ctx, sqlDB)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	// Common Interceptors
	commonInterceptors := connect.WithInterceptors(
		interseptor.NewCommonInterceptors()...,
	)

	// User Service
	mux.Handle(userv1connect.NewUserServiceHandler(
		handlers.UserServiceHandler,
		commonInterceptors,
	))

	return middleware.NewCommonMiddlewares(mux), nil
}
