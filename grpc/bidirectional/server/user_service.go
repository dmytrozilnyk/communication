package main

import (
	"context"
	userpb "github.com/dmytrozilnyk/communication/grpc/gen/go/proto/user/v1"
)

type userService struct {
	userpb.UnimplementedUserServiceServer
}

func (u *userService) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:     req.Uuid,
			FullName: "Dimi",
		},
	}, nil
}
