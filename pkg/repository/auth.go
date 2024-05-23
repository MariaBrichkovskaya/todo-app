package repository

import (
	"gorm.io/gorm"
	"todo-app/pkg/model"
)

type AuthRepository struct {
	db *gorm.DB
}

func (a *AuthRepository) CreateUser(request model.User) (*model.User, error) {

	user, _ := model.NewUser(
		request.Name, request.Username, request.Password,
	)
	result := a.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (a *AuthRepository) GetUser(username string, password string) model.User {

	var user model.User
	a.db.Where("username = ? AND password_hash = ?", username, password).Find(&user)

	return user
}
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
