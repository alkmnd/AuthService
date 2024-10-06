package main

import (
	"authService"
	handler2 "authService/pkg/handler"
	"authService/pkg/repository "
	service2 "authService/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "user",
		DBName:   "authservice",
		SSLMode:  "disable",
		Password: "qwerty",
	})
	if err != nil {
		logrus.Fatalf("error while connecting to the database: %s", err.Error())
	}
	repo := repository.NewAuthPostgres(db)
	service := service2.NewAuthService(repo)
	handler := handler2.NewHandler(service)
	srv := new(authService.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		logrus.Fatalf("error %s", err.Error())
	}
}
