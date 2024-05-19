package service

import (
	"firstGoProject/pkg/model"
	"firstGoProject/pkg/repository"
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
