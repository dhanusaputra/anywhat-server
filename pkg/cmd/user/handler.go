package user

import (
	"context"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Login ...
func (s *grpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	v := &loginRequest{
		Username: req.Username,
		Password: req.Password,
	}
	if err := s.validate.Struct(v); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := s.user.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token}, nil
}

// GetUser ...
func (s *grpcServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	u, err := s.user.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{User: u}, nil
}

// ListUser ...
func (s *grpcServer) ListUser(ctx context.Context, _ *emptypb.Empty) (*pb.ListUserResponse, error) {
	us, err := s.user.List(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListUserResponse{Users: us}, nil
}

// CreateUser ...
func (s *grpcServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if req.User == nil {
		return nil, status.Error(codes.InvalidArgument, "user empty")
	}

	v := &createUserRequest{
		Username: req.User.Username,
		Password: req.User.Password,
	}
	if err := s.validate.Struct(v); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := s.user.Create(ctx, req.User)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Id: id}, nil
}

// UpdateUser ...
func (s *grpcServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if req.User == nil {
		return nil, status.Error(codes.InvalidArgument, "user empty")
	}

	v := &updateUserRequest{
		Username: req.User.Username,
		Password: req.User.Password,
	}
	if err := s.validate.Struct(v); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	success, err := s.user.Update(ctx, req.User)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{Updated: success}, nil
}

// DeleteUser ...
func (s *grpcServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if req.Id == "1" {
		return nil, status.Error(codes.InvalidArgument, "cannot delete admin")
	}

	success, err := s.user.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Deleted: success}, nil
}
