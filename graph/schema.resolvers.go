package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ngoctd314/gql-stepbystep/graph/generated"
	"github.com/ngoctd314/gql-stepbystep/graph/model"
)

// Whoami is the resolver for the whoami field.
func (r *queryResolver) Whoami(ctx context.Context) (*model.Me, error) {
	me := &model.Me{
		ID:   "1",
		Name: "Gophers",
	}

	return me, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
