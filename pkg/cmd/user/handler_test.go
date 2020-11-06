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
			name: "failed login",
			args: args{
				ctx: ctx,
				req: &pb.LoginRequest{},
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
