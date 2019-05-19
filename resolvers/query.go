package resolvers

import (
	"context"

	gqlgen_todos "github.com/danishsatkut/gqlgen-todos"
	"github.com/danishsatkut/gqlgen-todos/models"
)

type queryResolver struct{
	store *gqlgen_todos.DataStore
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	return r.store.GetTodoList(), nil
}
