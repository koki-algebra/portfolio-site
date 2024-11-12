package repository

import (
	"context"

	"github.com/google/uuid"

	usermodel "backend/internal/domain/model/user"
)

type User interface {
	FindByID(ctx context.Context, userID uuid.UUID) (*usermodel.User, error)
	FindByEmail(ctx context.Context, email string) (*usermodel.User, error)
}
