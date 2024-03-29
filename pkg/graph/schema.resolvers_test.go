package graph

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/mocks"
	"github.com/dhanusaputra/anywhat-server/pkg/graph/model"
	"github.com/dhanusaputra/anywhat-server/util/authutil"
	"github.com/dhanusaputra/anywhat-server/util/testutil"
	"github.com/stretchr/testify/mock"
)

func TestCreateAnything(t *testing.T) {
	mockAnywhatClient := &mocks.AnywhatClient{}
	tests := []struct {
		name    string
		mock    func()
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockAnywhatClient.On("CreateAnything", mock.Anything, mock.Anything).Return(&pb.CreateAnythingResponse{Id: "mockID"}, nil).Once()
			},
			want: "mockID",
		},
		{
			name: "failed auth",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return nil
				}
			},
			wantErr: true,
		},
		{
			name: "failed create",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockAnywhatClient.On("CreateAnything", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&NewResolver, &authutil.GetUserContext}).Restore()
			tt.mock()
			r := NewResolver(mockAnywhatClient, nil)
			m := mutationResolver{r}
			got, err := m.CreateAnything(context.Background(), &model.AnythingInput{})
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
	mockAnywhatClient := &mocks.AnywhatClient{}
	tests := []struct {
		name    string
		mock    func()
		want    bool
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockAnywhatClient.On("UpdateAnything", mock.Anything, mock.Anything).Return(&pb.UpdateAnythingResponse{Updated: true}, nil).Once()
			},
			want: true,
		},
		{
			name: "failed auth",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return nil
				}
			},
			wantErr: true,
		},
		{
			name: "failed update",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockAnywhatClient.On("UpdateAnything", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&NewResolver, &authutil.GetUserContext}).Restore()
			tt.mock()
			r := NewResolver(mockAnywhatClient, nil)
			m := mutationResolver{r}
			got, err := m.UpdateAnything(context.Background(), "mockID", &model.AnythingInput{})
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
	mockAnywhatClient := &mocks.AnywhatClient{}
	tests := []struct {
		name    string
		mock    func()
		want    bool
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockAnywhatClient.On("DeleteAnything", mock.Anything, mock.Anything).Return(&pb.DeleteAnythingResponse{Deleted: true}, nil).Once()
			},
			want: true,
		},
		{
			name: "failed auth",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return nil
				}
			},
			wantErr: true,
		},
		{
			name: "failed delete",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockAnywhatClient.On("DeleteAnything", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&NewResolver, &authutil.GetUserContext}).Restore()
			tt.mock()
			r := NewResolver(mockAnywhatClient, nil)
			m := mutationResolver{r}
			got, err := m.DeleteAnything(context.Background(), "mockID")
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

func TestGetAnything(t *testing.T) {
	mockAnywhatClient := &mocks.AnywhatClient{}
	tests := []struct {
		name    string
		mock    func()
		want    *model.Anything
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				mockAnywhatClient.On("GetAnything", mock.Anything, mock.Anything).Return(&pb.GetAnythingResponse{Anything: &pb.Anything{}}, nil).Once()
			},
			want: &model.Anything{},
		},
		{
			name: "failed get",
			mock: func() {
				mockAnywhatClient.On("GetAnything", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&NewResolver}).Restore()
			tt.mock()
			r := NewResolver(mockAnywhatClient, nil)
			q := queryResolver{r}
			_, err := q.GetAnything(context.Background(), "mockID")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAnything() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestListAnything(t *testing.T) {
	mockAnywhatClient := &mocks.AnywhatClient{}
	tests := []struct {
		name    string
		mock    func()
		want    *model.Anything
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				mockAnywhatClient.On("ListAnything", mock.Anything, mock.Anything).Return(&pb.ListAnythingResponse{Anythings: []*pb.Anything{{}}}, nil).Once()
			},
			want: &model.Anything{},
		},
		{
			name: "failed list",
			mock: func() {
				mockAnywhatClient.On("ListAnything", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&NewResolver}).Restore()
			tt.mock()
			r := NewResolver(mockAnywhatClient, nil)
			q := queryResolver{r}
			_, err := q.ListAnything(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAnything() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLogin(t *testing.T) {
	mockUserClient := &mocks.UserServiceClient{}
	tests := []struct {
		name    string
		mock    func()
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				mockUserClient.On("Login", mock.Anything, mock.Anything).Return(&pb.LoginResponse{Token: ""}, nil).Once()
			},
			want: "",
		},
		{
			name: "failed login",
			mock: func() {
				mockUserClient.On("Login", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&NewResolver}).Restore()
			tt.mock()
			r := NewResolver(nil, mockUserClient)
			q := queryResolver{r}
			got, err := q.Login(context.Background(), "mockUsername", "mockPassword")
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
	mockUserClient := &mocks.UserServiceClient{}
	tests := []struct {
		name    string
		mock    func()
		want    *pb.User
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
			},
			want: &pb.User{},
		},
		{
			name: "failed me",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return nil
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&authutil.GetUserContext}).Restore()
			tt.mock()
			r := NewResolver(nil, mockUserClient)
			q := queryResolver{r}
			_, err := q.Me(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Me() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	mockUserClient := &mocks.UserServiceClient{}
	tests := []struct {
		name    string
		mock    func()
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockUserClient.On("CreateUser", mock.Anything, mock.Anything).Return(&pb.CreateUserResponse{Id: "mockID"}, nil).Once()
			},
			want: "mockID",
		},
		{
			name: "failed auth",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return nil
				}
			},
			wantErr: true,
		},
		{
			name: "failed create",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockUserClient.On("CreateUser", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&mockUserClient, &authutil.GetUserContext}).Restore()
			tt.mock()
			r := NewResolver(nil, mockUserClient)
			m := mutationResolver{r}
			got, err := m.CreateUser(context.Background(), &model.UserInput{})
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
	mockUserClient := &mocks.UserServiceClient{}
	tests := []struct {
		name    string
		mock    func()
		want    bool
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockUserClient.On("UpdateUser", mock.Anything, mock.Anything).Return(&pb.UpdateUserResponse{Updated: true}, nil).Once()
			},
			want: true,
		},
		{
			name: "failed auth",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return nil
				}
			},
			wantErr: true,
		},
		{
			name: "failed update",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockUserClient.On("UpdateUser", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&mockUserClient, &authutil.GetUserContext}).Restore()
			tt.mock()
			r := NewResolver(nil, mockUserClient)
			m := mutationResolver{r}
			got, err := m.UpdateUser(context.Background(), "mockID", &model.UserInput{})
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
	mockUserClient := &mocks.UserServiceClient{}
	tests := []struct {
		name    string
		mock    func()
		want    bool
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockUserClient.On("DeleteUser", mock.Anything, mock.Anything).Return(&pb.DeleteUserResponse{Deleted: true}, nil).Once()
			},
			want: true,
		},
		{
			name: "failed auth",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return nil
				}
			},
			wantErr: true,
		},
		{
			name: "failed delete",
			mock: func() {
				authutil.GetUserContext = func(ctx context.Context) *pb.User {
					return &pb.User{}
				}
				mockUserClient.On("DeleteUser", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&mockUserClient, &authutil.GetUserContext}).Restore()
			tt.mock()
			r := NewResolver(nil, mockUserClient)
			m := mutationResolver{r}
			got, err := m.DeleteUser(context.Background(), "mockID")
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

func TestGetUser(t *testing.T) {
	mockUserClient := &mocks.UserServiceClient{}
	tests := []struct {
		name    string
		mock    func()
		want    *model.User
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				mockUserClient.On("GetUser", mock.Anything, mock.Anything).Return(&pb.GetUserResponse{User: &pb.User{}}, nil).Once()
			},
			want: &model.User{},
		},
		{
			name: "failed get",
			mock: func() {
				mockUserClient.On("GetUser", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&mockUserClient}).Restore()
			tt.mock()
			r := NewResolver(nil, mockUserClient)
			q := queryResolver{r}
			_, err := q.GetUser(context.Background(), "mockID")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestListUser(t *testing.T) {
	mockUserClient := &mocks.UserServiceClient{}
	tests := []struct {
		name    string
		mock    func()
		want    *model.User
		wantErr bool
	}{
		{
			name: "happy path",
			mock: func() {
				mockUserClient.On("ListUser", mock.Anything, mock.Anything).Return(&pb.ListUserResponse{Users: []*pb.User{{}}}, nil).Once()
			},
			want: &model.User{},
		},
		{
			name: "failed list",
			mock: func() {
				mockUserClient.On("ListUser", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&mockUserClient}).Restore()
			tt.mock()
			r := NewResolver(nil, mockUserClient)
			q := queryResolver{r}
			_, err := q.ListUser(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
