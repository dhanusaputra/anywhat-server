package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/graph/generated"
	"github.com/dhanusaputra/anywhat-server/pkg/graph/model"
	"github.com/dhanusaputra/anywhat-server/util/authutil"
	"github.com/golang/protobuf/ptypes/empty"
)

func (r *mutationResolver) CreateAnything(ctx context.Context, input *model.AnythingInput) (string, error) {
	user := authutil.GetUserContext(ctx)
	if user == nil {
		return "", errors.New("access denied")
	}
	res, err := r.anywhatClient.CreateAnything(ctx, &pb.CreateAnythingRequest{Anything: &pb.Anything{
		Name:        input.Name,
		Description: input.Description,
	}})
	if err != nil {
		return "", err
	}
	return res.Id, nil
}

func (r *mutationResolver) UpdateAnything(ctx context.Context, id string, input *model.AnythingInput) (bool, error) {
	user := authutil.GetUserContext(ctx)
	if user == nil {
		return false, errors.New("access denied")
	}
	res, err := r.anywhatClient.UpdateAnything(ctx, &pb.UpdateAnythingRequest{Anything: &pb.Anything{
		Id:          id,
		Name:        input.Name,
		Description: input.Description,
	}})
	if err != nil {
		return false, err
	}
	return res.Updated, nil
}

func (r *mutationResolver) DeleteAnything(ctx context.Context, id string) (bool, error) {
	user := authutil.GetUserContext(ctx)
	if user == nil {
		return false, errors.New("access denied")
	}
	res, err := r.anywhatClient.DeleteAnything(ctx, &pb.DeleteAnythingRequest{Id: id})
	if err != nil {
		return false, err
	}
	return res.Deleted, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.UserInput) (string, error) {
	user := authutil.GetUserContext(ctx)
	if user == nil {
		return "", errors.New("access denied")
	}
	res, err := r.userClient.CreateUser(ctx, &pb.CreateUserRequest{User: &pb.User{
		Username: input.Username,
		Password: input.Password,
	}})
	if err != nil {
		return "", err
	}
	return res.Id, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.UserInput) (bool, error) {
	user := authutil.GetUserContext(ctx)
	if user == nil {
		return false, errors.New("access denied")
	}
	res, err := r.userClient.UpdateUser(ctx, &pb.UpdateUserRequest{User: &pb.User{
		Id:       id,
		Username: input.Username,
		Password: input.Password,
	}})
	if err != nil {
		return false, err
	}
	return res.Updated, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	user := authutil.GetUserContext(ctx)
	if user == nil {
		return false, errors.New("access denied")
	}
	res, err := r.userClient.DeleteUser(ctx, &pb.DeleteUserRequest{Id: id})
	if err != nil {
		return false, err
	}
	return res.Deleted, nil
}

func (r *queryResolver) GetAnything(ctx context.Context, id string) (*model.Anything, error) {
	res, err := r.anywhatClient.GetAnything(ctx, &pb.GetAnythingRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &model.Anything{
		ID:          res.Anything.Id,
		Name:        res.Anything.Name,
		Description: res.Anything.Description,
		CreatedAt:   res.Anything.CreatedAt.AsTime(),
		UpdatedAt:   res.Anything.UpdatedAt.AsTime(),
	}, nil
}

func (r *queryResolver) ListAnything(ctx context.Context) ([]*model.Anything, error) {
	res, err := r.anywhatClient.ListAnything(ctx, new(empty.Empty))
	if err != nil {
		return nil, err
	}
	a := res.Anythings
	m := make([]*model.Anything, 0, len(a))
	for _, aa := range a {
		m = append(m, &model.Anything{
			ID:          aa.Id,
			Name:        aa.Name,
			Description: aa.Description,
			CreatedAt:   aa.CreatedAt.AsTime(),
			UpdatedAt:   aa.UpdatedAt.AsTime(),
		})
	}
	return m, nil
}

func (r *queryResolver) Login(ctx context.Context, username string, password string) (string, error) {
	res, err := r.userClient.Login(ctx, &pb.LoginRequest{Username: username, Password: password})
	if err != nil {
		return "", err
	}
	return res.Token, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user := authutil.GetUserContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}
	return &model.User{
		ID:       user.Id,
		Username: user.Username,
	}, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	res, err := r.userClient.GetUser(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:          res.User.Id,
		Username:    res.User.Username,
		CreatedAt:   res.User.CreatedAt.AsTime(),
		UpdatedAt:   res.User.UpdatedAt.AsTime(),
		LastLoginAt: res.User.LastLoginAt.AsTime(),
	}, nil
}

func (r *queryResolver) ListUser(ctx context.Context) ([]*model.User, error) {
	res, err := r.userClient.ListUser(ctx, new(empty.Empty))
	if err != nil {
		return nil, err
	}
	a := res.Users
	m := make([]*model.User, 0, len(a))
	for _, aa := range a {
		m = append(m, &model.User{
			ID:          aa.Id,
			Username:    aa.Username,
			CreatedAt:   aa.CreatedAt.AsTime(),
			UpdatedAt:   aa.UpdatedAt.AsTime(),
			LastLoginAt: aa.LastLoginAt.AsTime(),
		})
	}
	return m, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
