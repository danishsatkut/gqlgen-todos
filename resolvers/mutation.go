package resolvers

import (
	"context"

	"github.com/danishsatkut/gqlgen-todos/models"
)

type mutationResolver struct{
	store *models.DataStore
}

func (r *mutationResolver) CompleteTodo(ctx context.Context, id string) (*models.Todo, error) {
	todo, err := r.store.FindTodo(id)
	if err != nil {
		return nil, err
	}

	todo.Done = true

	return r.store.UpdateTodo(todo)
}

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*models.User, error) {
	return r.store.CreateUser(name), nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, text string, user string) (*models.Todo, error) {
	return r.store.CreateTodo(text, user)
}

