package usecase

import (
	"context"

	"connectrpc.com/connect"
	"github.com/google/uuid"

	usermodel "backend/internal/domain/model/user"
	"backend/internal/domain/repository"
)

type User interface {
	GetByID(ctx context.Context, userID uuid.UUID) (*usermodel.User, error)
}

type userImpl struct {
	repo repository.User
}

func NewUser(repo repository.User) User {
	return &userImpl{
		repo: repo,
	}
}

func (u *userImpl) GetByID(
	ctx context.Context,
	userID uuid.UUID,
) (*usermodel.User, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return user, nil
}
