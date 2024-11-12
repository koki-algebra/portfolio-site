package usermodel

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	userID := uuid.New()
	authID := uuid.NewString()
	email := "example@example.com"

	user := &User{
		UserID: userID,
		AuthID: authID,
		Email:  email,
	}

	baseCtx := context.Background()

	ctx := ContextWithUser(baseCtx, user)

	got, ok := UserFromContext(ctx)
	assert.True(t, ok)
	assert.Equal(t, user, got)
}
