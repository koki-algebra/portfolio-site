package userservice

import (
	"backend/internal/domain/model"
	userv1 "backend/pkg/grpc/gen/user/v1"
)

func convertUser(user *model.User) *userv1.User {
	return &userv1.User{
		UserId: user.UserID.String(),
		Email:  user.Email,
	}
}
