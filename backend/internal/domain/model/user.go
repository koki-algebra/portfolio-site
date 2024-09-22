package model

import (
	"context"

	"github.com/google/uuid"
)

type ctxUserKey struct{}

type User struct {
	UserID uuid.UUID
	Email  string
}

func ContextWithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, ctxUserKey{}, user)
}

func UserFromContext(ctx context.Context) (*User, bool) {
	user, ok := ctx.Value(ctxUserKey{}).(*User)
	return user, ok
}
