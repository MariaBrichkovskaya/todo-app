package service

import (
	"firstGoProject/pkg/model"
	"firstGoProject/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (*model.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}
type TodoService struct {
	Authorization
	TodoList
	TodoItem
}

func NewTodoService(repos *repository.TodoRepository) *TodoService {
	return &TodoService{
		Authorization: NewAuthService(repos.Authorization),
	}
}
