package resolvers

import (
	"context"

	"github.com/danishsatkut/gqlgen-todos/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	return r.todos, nil
}
