package user

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/mocks"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.LoginRequest
	}
	var s *grpcServer
	var mockUser *mocks.User
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
				req: &pb.LoginRequest{},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("mockToken", nil)
				s = &grpcServer{mockUser}
			},
			want: &pb.LoginResponse{Token: "mockToken"},
		},
		{
			name: "failed login",
			args: args{
				ctx: ctx,
				req: &pb.LoginRequest{},
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("err"))
				s = &grpcServer{mockUser}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

func TestMe(t *testing.T) {
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
		want    *pb.MeResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Me", mock.Anything).Return(&pb.User{}, nil)
				s = &grpcServer{mockUser}
			},
			want: &pb.MeResponse{
				User: &pb.User{},
			},
		},
		{
			name: "failed me",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockUser = &mocks.User{}
				mockUser.On("Me", mock.Anything).Return(&pb.User{}, errors.New("err"))
				s = &grpcServer{mockUser}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := s.Me(tt.args.ctx, new(empty.Empty))
			if (err != nil) != tt.wantErr {
				t.Errorf("Me() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Me() = %v, want %v", got, tt.want)
			}
		})
	}
}
