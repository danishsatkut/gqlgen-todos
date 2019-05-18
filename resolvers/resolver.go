package resolvers

import (
	"github.com/danishsatkut/gqlgen-todos"
	"github.com/danishsatkut/gqlgen-todos/models"
	"github.com/danishsatkut/gqlgen-todos/resolvers/notes"
)

type Resolver struct{
	todos []models.Todo
}

func (r *Resolver) Todo() gqlgen_todos.TodoResolver {
	return &notes.TodoResolver{}
}

func (r *Resolver) Query() gqlgen_todos.QueryResolver {
	r.todos = todos

	return &queryResolver{r}
}

var todos = []models.Todo{
	{ID: "1", Text: "First", Done: false, UserID: "1"},
	{ID: "2", Text: "Second", Done: false, UserID: "2"},
	{ID: "3", Text: "Third", Done: false, UserID: "1"},
}
