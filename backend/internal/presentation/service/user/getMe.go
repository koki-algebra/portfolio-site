package userservice

import (
	"context"

	"connectrpc.com/connect"

	userv1 "backend/pkg/grpc/gen/user/v1"
)

func (s *userServiceImpl) GetMe(
	ctx context.Context,
	req *connect.Request[userv1.GetMeRequest],
) (*connect.Response[userv1.GetMeResponse], error) {
	resp := connect.NewResponse(&userv1.GetMeResponse{
		Id:    "3b9f87cf-12f5-4f29-a7f5-561f7d08f57d",
		Email: "example@example.com",
	})

	return resp, nil
}
