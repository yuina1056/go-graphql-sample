package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	graphql1 "test/domain/models/graphql"
	graphql2 "test/infrastructure/graphql"
)

type Resolver struct{}

func (r *mutationResolver) CreateTodo(ctx context.Context, input graphql1.NewTodo) (*graphql1.Todo, error) {
	panic("not implemented")
}

func (r *queryResolver) Todos(ctx context.Context) ([]*graphql1.Todo, error) {
	panic("not implemented")
}

// Mutation returns graphql2.MutationResolver implementation.
func (r *Resolver) Mutation() graphql2.MutationResolver { return &mutationResolver{r} }

// Query returns graphql2.QueryResolver implementation.
func (r *Resolver) Query() graphql2.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
