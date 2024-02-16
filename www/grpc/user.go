package grpc

import (
	"context"
	"log"

	sternx "sternx/www/grpc/gen/go"
)

type userServer struct {
	*Server
}

func newUserServer(server *Server) *userServer {
	return &userServer{
		Server: server,
	}
}

func (s *userServer) CreateUser(ctx context.Context, request *sternx.CreateUserRequest) (*sternx.UserResponse, error) {
	log.Println(request)
	// TODO implement me
	panic("implement me")
}

func (s *userServer) GetUserByID(ctx context.Context, request *sternx.GetUserRequest) (*sternx.UserResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s *userServer) GetUsers(ctx context.Context, request *sternx.GetUsersRequest) (*sternx.GetUsersResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s *userServer) UpdateUser(ctx context.Context, request *sternx.UpdateUserRequest) (*sternx.UserResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s *userServer) DeleteUser(ctx context.Context, request *sternx.DeleteUserRequest) (*sternx.DeleteUserResponse,
	error,
) {
	// TODO implement me
	panic("implement me")
}
