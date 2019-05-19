package gqlgen_todos

import (
	"errors"

	"github.com/google/uuid"

	"github.com/danishsatkut/gqlgen-todos/models"
)

type DataStore struct {
	todos []*models.Todo
	users []*models.User
}

func NewDataStore() *DataStore {
	todos := make([]*models.Todo, 0, 10)
	users := make([]*models.User, 0, 10)

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

	s.todos = append(s.todos, todo)

	return todo, nil
}

func (s *DataStore) CreateUser(name string) *models.User {
	user := &models.User{
		ID: uuid.New().String(),
		Name: name,
	}

	s.users = append(s.users, user)

	return user
}

func (s *DataStore) FindUser(id string) (*models.User, error) {
	var user *models.User

	for _, u := range s.users {
		if u.ID == id {
			user = u
		}
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *DataStore) GetTodos() []*models.Todo {
	return s.todos
}
