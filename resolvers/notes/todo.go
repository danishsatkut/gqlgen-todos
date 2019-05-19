package notes

import (
	"context"

	"github.com/danishsatkut/gqlgen-todos/models"
)

type TodoResolver struct{
	store *models.DataStore
}

func NewTodoResolver(store *models.DataStore) *TodoResolver {
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
