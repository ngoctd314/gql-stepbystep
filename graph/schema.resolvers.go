package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ngoctd314/gql-stepbystep/graph/generated"
	"github.com/ngoctd314/gql-stepbystep/graph/model"
)

// CreateHero is the resolver for the createHero field.
func (r *mutationResolver) CreateHero(ctx context.Context, hero model.NewHero) (*model.Character, error) {
	newHero := &model.Character{
		ID:   hero.ID,
		Name: hero.Name,
	}
	r.Storage[hero.ID] = newHero
	return newHero, nil
}

// Hero is the resolver for the hero field.
func (r *queryResolver) Hero(ctx context.Context, id string) (*model.Character, error) {
	return r.Storage[id], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
