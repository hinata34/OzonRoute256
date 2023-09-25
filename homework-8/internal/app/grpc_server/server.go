package grpcserver

import (
	"context"
	"homework-8/internal/app/grpc_server/pb"
	"homework-8/internal/app/user"
)

type Implementation struct {
	pb.UnimplementedGRPCServiceServer

	userRepo user.UserRepo
}

func NewImplementation(userRepo user.UserRepo) *Implementation {
	return &Implementation{userRepo: userRepo}
}

func (s *Implementation) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	id, err := s.userRepo.Add(ctx, &user.User{Name: in.User.Name, Age: in.User.Age})

	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{Id: id}, nil
}
