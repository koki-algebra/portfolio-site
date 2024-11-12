package usermodel

import "context"

type ctxUserKey struct{}

func ContextWithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, ctxUserKey{}, user)
}

func UserFromContext(ctx context.Context) (*User, bool) {
	user, ok := ctx.Value(ctxUserKey{}).(*User)
	return user, ok
}
