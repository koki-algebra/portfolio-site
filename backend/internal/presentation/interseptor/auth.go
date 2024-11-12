package interseptor

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"

	"backend/internal/config"
	usermodel "backend/internal/domain/model/user"
	"backend/internal/domain/repository"
)

type AuthInterceptor struct {
	repo repository.User
}

func NewAuthInterceptor(repo repository.User) *AuthInterceptor {
	return &AuthInterceptor{
		repo: repo,
	}
}

func (a *AuthInterceptor) Auth() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			rawTokenString := req.Header().Get("Authorization")
			if rawTokenString == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("id token is required"))
			}

			client, err := config.FirebaseApp.Auth(ctx)
			if err != nil {
				return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to initialize firebase auth client"))
			}

			token, err := client.VerifyIDToken(ctx, rawTokenString)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("unauthenticated user"))
			}

			user, err := a.repo.FindByAuthID(ctx, token.UID)
			if err != nil {
				if errors.Is(err, repository.ErrNotFound) {
					return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("user not found"))
				}

				return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to find user: %v", err))
			}

			ctx = usermodel.ContextWithUser(ctx, user)

			return next(ctx, req)
		}
	}
}
