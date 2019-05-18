package resolvers

import (
	"context"

	gqlgen_todos "github.com/danishsatkut/gqlgen-todos"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]gqlgen_todos.Todo, error) {
	return r.todos, nil
}
