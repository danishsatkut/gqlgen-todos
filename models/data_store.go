package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type DataStore struct {
	todos map[string]*Todo
	users map[string]*User
}

func NewDataStore() *DataStore {
	todos := make(map[string]*Todo)
	users := make(map[string]*User)

	return &DataStore{todos, users}
}

func (s *DataStore) CreateTodo(text string, userId string) (*Todo, error) {
	user, err := s.FindUser(userId)
	if err != nil {
		return nil, err
	}

	todo := &Todo{
		ID: uuid.New().String(),
		Text: text,
		UserID: user.ID,
		Done: false,
	}

	s.todos[todo.ID] = todo

	return todo, nil
}

func (s *DataStore) CreateUser(name string) *User {
	user := &User{
		ID: uuid.New().String(),
		Name: name,
	}

	s.users[user.ID] = user

	return user
}

func (s *DataStore) FindUser(id string) (*User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *DataStore) GetTodoList() []*Todo {
	todoList := make([]*Todo, 0, len(s.todos))

	for _, todo := range s.todos {
		todoList = append(todoList, todo)
	}

	return todoList
}

func (s *DataStore) FindTodo(id string) (*Todo, error) {
	todo, ok := s.todos[id]
	if !ok {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

func (s *DataStore) UpdateTodo(todo *Todo) (*Todo, error) {
	if s.todos[todo.ID] == nil {
		return nil, fmt.Errorf("todo does not exist with id: %v", todo.ID)
	}

	s.todos[todo.ID] = todo

	return todo, nil
}
