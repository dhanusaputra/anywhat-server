package user

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/mocks"
	"github.com/dhanusaputra/anywhat-server/util/testutil"
	"github.com/go-playground/validator"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.LoginRequest
	}
	var (
		s        *grpcServer
		mockUser *mocks.User
		v        *validator.Validate = validator.New()
	)
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.LoginResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &pb.LoginRequest{Username: "mock", Password: "mock"},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("mockToken", nil)
				s = &grpcServer{mockUser, v}
			},
			want: &pb.LoginResponse{Token: "mockToken"},
		},
		{
			name: "failed validate",
			args: args{
				ctx: ctx,
				req: &pb.LoginRequest{Password: "m"},
			},
			mock: func() {
				s = &grpcServer{mockUser, v}
			},
			wantErr: true,
		},
		{
			name: "failed login",
			args: args{
				ctx: ctx,
				req: &pb.LoginRequest{Username: "mock", Password: "mock"},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("err"))
				s = &grpcServer{mockUser, v}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&s, &mockUser}).Restore()
			tt.mock()
			got, err := s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.GetUserRequest
	}
	var s *grpcServer
	var mockUser *mocks.User
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.GetUserResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &pb.GetUserRequest{},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Get", mock.Anything, mock.Anything).Return(&pb.User{}, nil)
			},
			want: &pb.GetUserResponse{
				User: &pb.User{},
			},
		},
		{
			name: "failed get",
			args: args{
				ctx: ctx,
				req: &pb.GetUserRequest{},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New("err"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&mockUser}).Restore()
			tt.mock()
			s = &grpcServer{mockUser, nil}
			got, err := s.GetUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListUser(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
	}
	var s *grpcServer
	var mockUser *mocks.User
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.ListUserResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("List", mock.Anything, mock.Anything).Return([]*pb.User{}, nil)
			},
			want: &pb.ListUserResponse{
				Users: []*pb.User{},
			},
		},
		{
			name: "failed list",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("List", mock.Anything, mock.Anything).Return(nil, errors.New("err"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&mockUser}).Restore()
			tt.mock()
			s = &grpcServer{mockUser, nil}
			got, err := s.ListUser(tt.args.ctx, new(empty.Empty))
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.CreateUserRequest
	}
	var (
		s        *grpcServer
		mockUser *mocks.User
		v        *validator.Validate = validator.New()
	)
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.CreateUserResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &pb.CreateUserRequest{User: &pb.User{Username: "mock", Password: "mock"}},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Create", mock.Anything, mock.Anything).Return("", nil)
			},
			want: &pb.CreateUserResponse{
				Id: "",
			},
		},
		{
			name: "anything nill",
			args: args{
				ctx: ctx,
				req: &pb.CreateUserRequest{},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "failed validate",
			args: args{
				ctx: ctx,
				req: &pb.CreateUserRequest{User: &pb.User{Username: "m"}},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "failed create",
			args: args{
				ctx: ctx,
				req: &pb.CreateUserRequest{User: &pb.User{Username: "mock", Password: "mock"}},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Create", mock.Anything, mock.Anything).Return("", errors.New("err"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&mockUser}).Restore()
			tt.mock()
			s = &grpcServer{mockUser, v}
			got, err := s.CreateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.UpdateUserRequest
	}
	var (
		s        *grpcServer
		mockUser *mocks.User
		v        *validator.Validate = validator.New()
	)
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.UpdateUserResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &pb.UpdateUserRequest{User: &pb.User{Password: "mock"}},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Update", mock.Anything, mock.Anything).Return(true, nil)
			},
			want: &pb.UpdateUserResponse{
				Updated: true,
			},
		},
		{
			name: "anything nil",
			args: args{
				ctx: ctx,
				req: &pb.UpdateUserRequest{},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "failed validate",
			args: args{
				ctx: ctx,
				req: &pb.UpdateUserRequest{User: &pb.User{Username: "m"}},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "failed update",
			args: args{
				ctx: ctx,
				req: &pb.UpdateUserRequest{User: &pb.User{Password: "mock"}},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Update", mock.Anything, mock.Anything).Return(false, errors.New("err"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&mockUser}).Restore()
			tt.mock()
			s = &grpcServer{mockUser, v}
			got, err := s.UpdateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.DeleteUserRequest
	}
	var s *grpcServer
	var mockUser *mocks.User
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.DeleteUserResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &pb.DeleteUserRequest{},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Delete", mock.Anything, mock.Anything).Return(true, nil)
			},
			want: &pb.DeleteUserResponse{
				Deleted: true,
			},
		},
		{
			name: "failed validate",
			args: args{
				ctx: ctx,
				req: &pb.DeleteUserRequest{Id: "1"},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "failed delete",
			args: args{
				ctx: ctx,
				req: &pb.DeleteUserRequest{},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Delete", mock.Anything, mock.Anything).Return(false, errors.New("err"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&mockUser}).Restore()
			tt.mock()
			s = &grpcServer{mockUser, nil}
			got, err := s.DeleteUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
