package notes

import (
	"context"

	gqlgen_todos "github.com/danishsatkut/gqlgen-todos"
	"github.com/danishsatkut/gqlgen-todos/models"
)

type TodoResolver struct{
	store *gqlgen_todos.DataStore
}

func NewTodoResolver(store *gqlgen_todos.DataStore) *TodoResolver {
	return &TodoResolver{store: store}
}

func (r *TodoResolver) Completed(ctx context.Context, obj *models.Todo) (bool, error) {
	return obj.Done, nil
}

func (r *TodoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	user, err := r.store.FindUser(obj.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
