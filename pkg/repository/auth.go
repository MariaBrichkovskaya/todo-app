package repository

import (
	"firstGoProject/pkg/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func (a *AuthRepository) CreateUser(request model.User) (*model.User, error) {

	user, _ := model.NewUser(
		request.Name, request.Username, request.PasswordHash,
	)
	result := a.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
