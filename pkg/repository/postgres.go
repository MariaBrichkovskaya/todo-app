package repository

import (
	"firstGoProject/pkg/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Password string
	Username string
	Database string
	SSLMode  string
}

func NewPostgresDB(cnf Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.Host, cnf.Port, cnf.Username, cnf.Password, cnf.Database, cnf.SSLMode)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
