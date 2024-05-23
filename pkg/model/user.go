package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"ID" gorm:"primaryKey" gorm:"column:id"`
	Name     string    `json:"name" binding:"required"`
	Password string    `json:"password" binding:"required" gorm:"column:password_hash"`
	Username string    `json:"username" binding:"required" gorm:"unique"`
}

func NewUser(name, username, password string) (*User, error) {
	id := uuid.New()
	return &User{
		ID:       id,
		Name:     name,
		Password: password,
		Username: username,
	}, nil
}
