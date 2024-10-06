package main

import (
	"authService"
	handler2 "authService/pkg/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	handler := handler2.NewHandler()
	srv := new(authService.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		logrus.Fatalf("error %s", err.Error())
	}
}
