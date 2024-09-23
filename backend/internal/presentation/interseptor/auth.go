package interseptor

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"google.golang.org/api/idtoken"

	"backend/internal/config"
	"backend/internal/domain/model"
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
			resp, err := next(ctx, req)
			if err != nil {
				return resp, err
			}

			rawTokenString := req.Header().Get("Authorization")
			if rawTokenString == "" {
				return resp, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("id token is required"))
			}

			audience := config.Env.GoogleClientID

			payload, err := idtoken.Validate(ctx, rawTokenString, audience)
			if err != nil {
				return resp, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid id token: %v", err))
			}

			email, ok := a.getEmail(payload.Claims)
			if !ok {
				return resp, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("the id token does not contain an email"))
			}

			user, err := a.repo.FindByEmail(ctx, email)
			if err != nil {
				if errors.Is(err, repository.ErrNotFound) {
					return resp, connect.NewError(connect.CodeNotFound, fmt.Errorf("user not found"))
				}

				return resp, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to find user: %v", err))
			}

			ctx = model.ContextWithUser(ctx, user)

			return next(ctx, req)
		}
	}
}

func (a *AuthInterceptor) getEmail(claims map[string]any) (string, bool) {
	v, ok := claims["email"]
	if !ok {
		return "", false
	}

	email, ok := v.(string)
	if !ok {
		return "", false
	}

	return email, true
}
