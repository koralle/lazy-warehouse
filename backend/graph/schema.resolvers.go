package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"

	"github.com/koralle/lazy-warehouse/backend/graph/model"
	"github.com/koralle/lazy-warehouse/backend/internal"
)

// GetAllAvailableRoles is the resolver for the getAllAvailableRoles field.
func (r *queryResolver) GetAllAvailableRoles(ctx context.Context) ([]*model.Role, error) {
	return r.Srv.GetAllAvailableRole(ctx)
}

// Group is the resolver for the group field.
func (r *queryResolver) Group(ctx context.Context, id string) (*model.Group, error) {
	return r.Srv.Group(ctx, id)
}

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
