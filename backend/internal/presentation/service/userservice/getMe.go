package userservice

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	usermodel "backend/internal/domain/model/user"
	userv1 "backend/pkg/grpc/gen/user/v1"
)

func (s *userServiceImpl) GetMe(
	ctx context.Context,
	req *connect.Request[userv1.GetMeRequest],
) (*connect.Response[userv1.GetMeResponse], error) {
	user, ok := usermodel.UserFromContext(ctx)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("user not set in context"))
	}

	resp := connect.NewResponse(&userv1.GetMeResponse{
		User: convertUser(user),
	})

	return resp, nil
}
