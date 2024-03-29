package anywhat

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

func TestGetAnything(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.GetAnythingRequest
	}
	var s *grpcServer
	var mockAnywhat *mocks.Anywhat
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.GetAnythingResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &pb.GetAnythingRequest{},
			},
			mock: func() {
				mockAnywhat = &mocks.Anywhat{}
				mockAnywhat.On("Get", mock.Anything, mock.Anything).Return(&pb.Anything{}, nil)
				s = &grpcServer{mockAnywhat, nil}
			},
			want: &pb.GetAnythingResponse{
				Anything: &pb.Anything{},
			},
		},
		{
			name: "failed get",
			args: args{
				ctx: ctx,
				req: &pb.GetAnythingRequest{},
			},
			mock: func() {
				mockAnywhat = &mocks.Anywhat{}
				mockAnywhat.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New("err"))
				s = &grpcServer{mockAnywhat, nil}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&s, &mockAnywhat}).Restore()
			tt.mock()
			got, err := s.GetAnything(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAnything() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAnything() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAnything(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
	}
	var s *grpcServer
	var mockAnywhat *mocks.Anywhat
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.ListAnythingResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockAnywhat = &mocks.Anywhat{}
				mockAnywhat.On("List", mock.Anything, mock.Anything).Return([]*pb.Anything{}, nil)
				s = &grpcServer{mockAnywhat, nil}
			},
			want: &pb.ListAnythingResponse{
				Anythings: []*pb.Anything{},
			},
		},
		{
			name: "failed list",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockAnywhat = &mocks.Anywhat{}
				mockAnywhat.On("List", mock.Anything, mock.Anything).Return(nil, errors.New("err"))
				s = &grpcServer{mockAnywhat, nil}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&s, &mockAnywhat}).Restore()
			tt.mock()
			got, err := s.ListAnything(tt.args.ctx, new(empty.Empty))
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAnything() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAnything() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateAnything(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.CreateAnythingRequest
	}
	var (
		s           *grpcServer
		mockAnywhat *mocks.Anywhat
		v           *validator.Validate = validator.New()
	)
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.CreateAnythingResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &pb.CreateAnythingRequest{Anything: &pb.Anything{Name: "mock"}},
			},
			mock: func() {
				mockAnywhat = &mocks.Anywhat{}
				mockAnywhat.On("Create", mock.Anything, mock.Anything).Return("", nil)
				s = &grpcServer{mockAnywhat, v}
			},
			want: &pb.CreateAnythingResponse{
				Id: "",
			},
		},
		{
			name: "anything nill",
			args: args{
				ctx: ctx,
				req: &pb.CreateAnythingRequest{},
			},
			mock: func() {
				s = &grpcServer{mockAnywhat, v}
			},
			wantErr: true,
		},
		{
			name: "failed validate",
			args: args{
				ctx: ctx,
				req: &pb.CreateAnythingRequest{Anything: &pb.Anything{Name: "m"}},
			},
			mock: func() {
				s = &grpcServer{mockAnywhat, v}
			},
			wantErr: true,
		},
		{
			name: "failed create",
			args: args{
				ctx: ctx,
				req: &pb.CreateAnythingRequest{Anything: &pb.Anything{Name: "mock"}},
			},
			mock: func() {
				mockAnywhat = &mocks.Anywhat{}
				mockAnywhat.On("Create", mock.Anything, mock.Anything).Return("", errors.New("err"))
				s = &grpcServer{mockAnywhat, v}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&s, &mockAnywhat}).Restore()
			tt.mock()
			got, err := s.CreateAnything(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAnything() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAnything() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateAnything(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.UpdateAnythingRequest
	}
	var (
		s           *grpcServer
		mockAnywhat *mocks.Anywhat
		v           *validator.Validate = validator.New()
	)
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.UpdateAnythingResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &pb.UpdateAnythingRequest{Anything: &pb.Anything{Name: "mock"}},
			},
			mock: func() {
				mockAnywhat = &mocks.Anywhat{}
				mockAnywhat.On("Update", mock.Anything, mock.Anything).Return(true, nil)
				s = &grpcServer{mockAnywhat, v}
			},
			want: &pb.UpdateAnythingResponse{
				Updated: true,
			},
		},
		{
			name: "anything nil",
			args: args{
				ctx: ctx,
				req: &pb.UpdateAnythingRequest{},
			},
			mock: func() {
				s = &grpcServer{mockAnywhat, v}
			},
			wantErr: true,
		},
		{
			name: "failed validate",
			args: args{
				ctx: ctx,
				req: &pb.UpdateAnythingRequest{Anything: &pb.Anything{Name: "m"}},
			},
			mock: func() {
				s = &grpcServer{mockAnywhat, v}
			},
			wantErr: true,
		},
		{
			name: "failed update",
			args: args{
				ctx: ctx,
				req: &pb.UpdateAnythingRequest{Anything: &pb.Anything{Name: "mock"}},
			},
			mock: func() {
				mockAnywhat = &mocks.Anywhat{}
				mockAnywhat.On("Update", mock.Anything, mock.Anything).Return(false, errors.New("err"))
				s = &grpcServer{mockAnywhat, v}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&s, &mockAnywhat}).Restore()
			tt.mock()
			got, err := s.UpdateAnything(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAnything() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAnything() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteAnything(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.DeleteAnythingRequest
	}
	var s *grpcServer
	var mockAnywhat *mocks.Anywhat
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *pb.DeleteAnythingResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &pb.DeleteAnythingRequest{},
			},
			mock: func() {
				mockAnywhat = &mocks.Anywhat{}
				mockAnywhat.On("Delete", mock.Anything, mock.Anything).Return(true, nil)
				s = &grpcServer{mockAnywhat, nil}
			},
			want: &pb.DeleteAnythingResponse{
				Deleted: true,
			},
		},
		{
			name: "failed delete",
			args: args{
				ctx: ctx,
				req: &pb.DeleteAnythingRequest{},
			},
			mock: func() {
				mockAnywhat = &mocks.Anywhat{}
				mockAnywhat.On("Delete", mock.Anything, mock.Anything).Return(false, errors.New("err"))
				s = &grpcServer{mockAnywhat, nil}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&s, &mockAnywhat}).Restore()
			tt.mock()
			got, err := s.DeleteAnything(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAnything() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAnything() = %v, want %v", got, tt.want)
			}
		})
	}
}
