// di/wire.go
//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"database/sql"

	"github.com/google/wire"

	"backend/internal/application/usecase"
	"backend/internal/infrastructure/database"
	"backend/internal/infrastructure/repository"
	"backend/internal/presentation/interseptor"
	"backend/internal/presentation/service/userservice"
	"backend/pkg/grpc/gen/user/v1/userv1connect"
)

var infrastructureSet = wire.NewSet(
	database.NewSqlxDB,
)

var repositorySet = wire.NewSet(
	repository.NewUser,
)

var usecaseSet = wire.NewSet(
	usecase.NewUser,
)

var connectServiceSet = wire.NewSet(
	userservice.New,
)

var interceptorSet = wire.NewSet(
	interseptor.NewAuthInterceptor,
)

type UsecaseSet struct {
	User usecase.User
}

type ConnectServiceSet struct {
	UserServiceHandler userv1connect.UserServiceHandler
}

type InterceptorSet struct {
	AuthInterceptor *interseptor.AuthInterceptor
}

func InitConnectService(ctx context.Context, db *sql.DB) (*ConnectServiceSet, error) {
	wire.Build(
		infrastructureSet,
		repositorySet,
		usecaseSet,
		connectServiceSet,
		wire.Struct(new(ConnectServiceSet), "*"),
	)

	return &ConnectServiceSet{}, nil
}

func InitInterceptor(ctx context.Context, db *sql.DB) (*InterceptorSet, error) {
	wire.Build(
		infrastructureSet,
		repositorySet,
		interceptorSet,
		wire.Struct(new(InterceptorSet), "*"),
	)

	return &InterceptorSet{}, nil
}
