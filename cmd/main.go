package main

import (
	"firstGoProject"
	"firstGoProject/pkg/handler"
	"firstGoProject/pkg/repository"
	"firstGoProject/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("Error init config ", err)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error load environment variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}
	repos := repository.NewTodoRepository(db)
	services := service.NewTodoService(repos)
	handlers := handler.NewHandler(services)

	svr := new(firstGoProject.Server)
	if err := svr.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal("error starting server:", err)
	}
}
func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/firstGoProject/configs/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
