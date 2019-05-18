package notes

import (
	"context"

	gqlgen_todos "github.com/danishsatkut/gqlgen-todos"
)

type TodoResolver struct{}

func (r *TodoResolver) User(ctx context.Context, obj *gqlgen_todos.Todo) (*gqlgen_todos.User, error) {
	return &gqlgen_todos.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
