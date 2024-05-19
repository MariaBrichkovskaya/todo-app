package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `json:"ID" gorm:"primaryKey" gorm:"column:id"`
	Name         string    `json:"name" binding:"required"`
	PasswordHash string    `json:"password" binding:"required" gorm:"column:password_hash"`
	Username     string    `json:"username" binding:"required" gorm:"unique"`
}

func NewUser(name, username, password string) (*User, error) {
	id := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:           id,
		Name:         name,
		PasswordHash: string(hashedPassword),
		Username:     username,
	}, nil
}
