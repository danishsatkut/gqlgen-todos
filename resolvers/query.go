package resolvers

import (
	"context"

	"github.com/danishsatkut/gqlgen-todos/models"
)

type queryResolver struct{
	store *models.DataStore
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	return r.store.GetTodoList(), nil
}
