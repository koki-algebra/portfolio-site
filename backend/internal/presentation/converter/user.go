package converter

import (
	usermodel "backend/internal/domain/model/user"
	userv1 "backend/pkg/grpc/gen/user/v1"
)

func ConvertUser(user *usermodel.User) *userv1.User {
	if user == nil {
		return nil
	}

	return &userv1.User{
		UserId: user.UserID.String(),
		AuthId: user.AuthID,
		Email:  user.Email,
	}
}
