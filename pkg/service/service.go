package service

import (
	"todo-app/pkg/model"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (*model.User, error)
	GenerateToken(username string, password string) (string, error)
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
