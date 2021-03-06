package resolvers

import (
	"github.com/danishsatkut/gqlgen-todos"
	"github.com/danishsatkut/gqlgen-todos/models"
	"github.com/danishsatkut/gqlgen-todos/resolvers/notes"
)

type Resolver struct{
	store *models.DataStore
}

func (r *Resolver) Todo() gqlgen_todos.TodoResolver {
	return notes.NewTodoResolver(r.store)
}

func (r *Resolver) Query() gqlgen_todos.QueryResolver {
	return &queryResolver{store: r.store}
}

func (r *Resolver) Mutation() gqlgen_todos.MutationResolver {
	return &mutationResolver{store: r.store}
}

func NewRootResolver(store *models.DataStore) *Resolver {
	return &Resolver{store}
}
