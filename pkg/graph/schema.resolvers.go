package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/graph/generated"
	"github.com/dhanusaputra/anywhat-server/pkg/graph/model"
)

func (r *mutationResolver) CreateAnything(ctx context.Context, input *model.AnythingInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAnything(ctx context.Context, id string, input *model.AnythingInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, username string, password string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
