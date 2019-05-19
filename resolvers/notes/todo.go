package notes

import (
	"context"

	"github.com/danishsatkut/gqlgen-todos/models"
)

type TodoResolver struct{}

func (r *TodoResolver) Completed(ctx context.Context, obj *models.Todo) (bool, error) {
	return obj.Done, nil
}

func (r *TodoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	return &models.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
