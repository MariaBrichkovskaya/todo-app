package repository

import (
	"gorm.io/gorm"
	"todo-app/pkg/model"
)

type Authorization interface {
	CreateUser(user model.User) (*model.User, error)
	GetUser(username string, password string) model.User
}

type TodoList interface {
}

type TodoItem interface {
}
type TodoRepository struct {
	Authorization
	TodoList
	TodoItem
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		Authorization: NewAuthRepository(db),
	}
}
