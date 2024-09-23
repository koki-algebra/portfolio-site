package userservice

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"backend/internal/domain/model"
	userv1 "backend/pkg/grpc/gen/user/v1"
)

func (s *userServiceImpl) GetMe(
	ctx context.Context,
	req *connect.Request[userv1.GetMeRequest],
) (*connect.Response[userv1.GetMeResponse], error) {
	user, ok := model.UserFromContext(ctx)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("user not set in context"))
	}

	resp := connect.NewResponse(&userv1.GetMeResponse{
		User: convertUser(user),
	})

	return resp, nil
}
