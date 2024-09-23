package interseptor

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

func UnexpectedErrorHandler() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err == nil {
				return resp, nil
			}

			if _, ok := err.(*connect.Error); !ok {
				slog.ErrorContext(ctx, "got unexpected error: error must be connect error", "error", err)
				return resp, connect.NewError(connect.CodeInternal, err)
			}

			return resp, err
		}
	}
}
