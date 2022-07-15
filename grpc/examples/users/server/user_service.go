package main

import (
	"context"
	"fmt"
	userpb "github.com/dmytrozilnyk/communication/grpc/gen/go/proto/user/v1"
)

type userService struct {
	userpb.UnimplementedUserServiceServer
}

func (u *userService) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	fmt.Println("Requested user:", req.Uuid)
	salary := uint32(1000)
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:          req.Uuid,
			FullName:      "Dimi",
			BirthYear:     191993,
			Salary:        &salary,
			MaritalStatus: 1,
		},
	}, nil
}
