package resolvers

import (
	"context"

	"github.com/danishsatkut/gqlgen-todos"
	"github.com/danishsatkut/gqlgen-todos/models"
)

type mutationResolver struct{
	store *gqlgen_todos.DataStore
}

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*models.User, error) {
	return r.store.CreateUser(name), nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, text string, user string) (*models.Todo, error) {
	return r.store.CreateTodo(text, user)
}

