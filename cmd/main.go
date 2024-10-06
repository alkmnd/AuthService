package main

import (
	"authService"
	handler2 "authService/pkg/handler"
	service2 "authService/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	service := service2.NewAuthService()
	handler := handler2.NewHandler(service)
	srv := new(authService.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		logrus.Fatalf("error %s", err.Error())
	}
}
