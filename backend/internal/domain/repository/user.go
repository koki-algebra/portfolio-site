package repository

import (
	"context"

	"github.com/google/uuid"

	"backend/internal/domain/model"
)

type User interface {
	FindByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
