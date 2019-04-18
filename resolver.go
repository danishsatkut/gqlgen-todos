package gqlgen_todos

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	todos []Todo
}

func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	r.todos = todos

	return &queryResolver{r}
}

var todos = []Todo{
	{ID: "1", Text: "First", Done: false},
	{ID: "2", Text: "Second", Done: false},
	{ID: "3", Text: "Third", Done: false},
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]Todo, error) {
	return r.todos, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
	return &User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
