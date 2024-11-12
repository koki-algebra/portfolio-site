package userservice

import (
	usermodel "backend/internal/domain/model/user"
	userv1 "backend/pkg/grpc/gen/user/v1"
)

func convertUser(user *usermodel.User) *userv1.User {
	return &userv1.User{
		UserId: user.UserID.String(),
		Email:  user.Email,
	}
}
