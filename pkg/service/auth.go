package service

import (
	"todo-app/pkg/model"
	"todo-app/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user model.User) (*model.User, error) {
	return s.repo.CreateUser(user)
}
