package gqlgen_todos

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/danishsatkut/gqlgen-todos/models"
)

type DataStore struct {
	todos map[string]*models.Todo
	users map[string]*models.User
}

func NewDataStore() *DataStore {
	todos := make(map[string]*models.Todo)
	users := make(map[string]*models.User)

	return &DataStore{todos, users}
}

func (s *DataStore) CreateTodo(text string, userId string) (*models.Todo, error) {
	user, err := s.FindUser(userId)
	if err != nil {
		return nil, err
	}

	todo := &models.Todo{
		ID: uuid.New().String(),
		Text: text,
		UserID: user.ID,
		Done: false,
	}

	s.todos[todo.ID] = todo

	return todo, nil
}

func (s *DataStore) CreateUser(name string) *models.User {
	user := &models.User{
		ID: uuid.New().String(),
		Name: name,
	}

	s.users[user.ID] = user

	return user
}

func (s *DataStore) FindUser(id string) (*models.User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *DataStore) GetTodoList() []*models.Todo {
	todoList := make([]*models.Todo, 0, len(s.todos))

	for _, todo := range s.todos {
		todoList = append(todoList, todo)
	}

	return todoList
}

func (s *DataStore) FindTodo(id string) (*models.Todo, error) {
	todo, ok := s.todos[id]
	if !ok {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

func (s *DataStore) UpdateTodo(todo *models.Todo) (*models.Todo, error) {
	if s.todos[todo.ID] == nil {
		return nil, fmt.Errorf("todo does not exist with id: %v", todo.ID)
	}

	s.todos[todo.ID] = todo

	return todo, nil
}
