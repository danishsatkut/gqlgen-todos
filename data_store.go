package gqlgen_todos

import "github.com/danishsatkut/gqlgen-todos/models"

type DataStore struct {
	todos []*models.Todo
}

func NewDataStore() *DataStore {
	todos := make([]*models.Todo, 0, 10)

	return &DataStore{todos}
}

func (s *DataStore) Add(todo *models.Todo) {
	s.todos = append(s.todos, todo)
}

func (s *DataStore) GetTodos() []*models.Todo {
	return s.todos
}
