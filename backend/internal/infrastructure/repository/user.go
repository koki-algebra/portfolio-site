package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"backend/internal/domain/model"
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

func (repo *userImpl) FindByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return &model.User{}, nil
}

func (repo *userImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return &model.User{}, nil
}
