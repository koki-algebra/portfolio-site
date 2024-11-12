package callermodel

import (
	"context"

	usermodel "backend/internal/domain/model/user"
)

type ctxCallerKey struct{}

func ContextWithCaller(ctx context.Context, user *usermodel.User) context.Context {
	if user == nil {
		return ctx
	}

	caller := &Caller{
		userID: user.UserID,
		authID: user.AuthID,
		email:  user.Email,
	}

	return context.WithValue(ctx, ctxCallerKey{}, caller)
}

func CallerFromContext(ctx context.Context) (*Caller, bool) {
	caller, ok := ctx.Value(ctxCallerKey{}).(*Caller)
	return caller, ok
}
