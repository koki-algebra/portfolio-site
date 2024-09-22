package userservice

import (
	"backend/internal/application/usecase"
	"backend/pkg/grpc/gen/user/v1/userv1connect"
)

type userServiceImpl struct {
	user usecase.User
}

func New(user usecase.User) userv1connect.UserServiceHandler {
	return &userServiceImpl{
		user: user,
	}
}
