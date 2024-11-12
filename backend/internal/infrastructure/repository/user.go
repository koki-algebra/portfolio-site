package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	usermodel "backend/internal/domain/model/user"
	"backend/internal/domain/repository"
)

type userImpl struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repository.User {
	return &userImpl{
		db: db,
	}
}

func (repo *userImpl) FindByID(ctx context.Context, userID uuid.UUID) (*usermodel.User, error) {
	return &usermodel.User{}, nil
}

func (repo *userImpl) FindByAuthID(ctx context.Context, email string) (*usermodel.User, error) {
	return &usermodel.User{}, nil
}
