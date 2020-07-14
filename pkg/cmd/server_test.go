package cmd

import (
	"context"
	"reflect"
	"testing"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/mocks"
)

func TestGetAnything(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *pb.GetAnythingRequest
	}
	s := &grpcServer{}
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
				mockAnywhat := &mocks.Anywhat{}
				mockAnywhat.On("Get", &pb.Anything{}, nil)
			},
			want: &pb.GetAnythingResponse{
				Anything: &pb.Anything{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
