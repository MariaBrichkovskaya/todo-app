package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
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
		Host:     os.Getenv("DB_HOST"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USERNAME"),
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

	svr := new(todo_app.Server)
	if err := svr.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal("error starting server:", err)
	}
}
func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
