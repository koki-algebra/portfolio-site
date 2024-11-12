package usermodel

import (
	"github.com/google/uuid"
)

type User struct {
	UserID uuid.UUID
	AuthID string
	Email  string
}
