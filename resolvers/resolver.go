package resolvers

import (
	"github.com/danishsatkut/gqlgen-todos"
	"github.com/danishsatkut/gqlgen-todos/resolvers/notes"
)

type Resolver struct{
	store *gqlgen_todos.DataStore
}

func (r *Resolver) Todo() gqlgen_todos.TodoResolver {
	return &notes.TodoResolver{}
}

func (r *Resolver) Query() gqlgen_todos.QueryResolver {
	return &queryResolver{store: r.store}
}

func NewRootResolver(store *gqlgen_todos.DataStore) *Resolver {
	return &Resolver{store}
}
