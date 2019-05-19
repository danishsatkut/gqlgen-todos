package resolvers

import (
	"context"

	"github.com/google/uuid"

	"github.com/danishsatkut/gqlgen-todos"
	"github.com/danishsatkut/gqlgen-todos/models"
)

type mutationResolver struct{
	store *gqlgen_todos.DataStore
}

func (r *mutationResolver) CreateTodo(ctx context.Context, text string, user string) (*models.Todo, error) {
	todo := &models.Todo{
		ID: uuid.New().String(),
		Text: text,
		UserID: user,
		Done: false,
	}

	r.store.Add(todo)

	return todo, nil
}

