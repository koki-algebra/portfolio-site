package callermodel

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	usermodel "backend/internal/domain/model/user"
)

func TestContextWithCaller(t *testing.T) {
	userID := uuid.New()
	authID := uuid.NewString()
	email := "example@example.com"

	user := &usermodel.User{
		UserID: userID,
		AuthID: authID,
		Email:  email,
	}

	baseCtx := context.Background()

	ContextWithCaller(baseCtx, nil)

	ctx := ContextWithCaller(baseCtx, user)

	caller, ok := CallerFromContext(ctx)
	assert.True(t, ok)

	assert.Equal(t, userID, caller.GetUserID())
	assert.Equal(t, authID, caller.GetAuthID())
	assert.Equal(t, email, caller.GetEmail())
}
