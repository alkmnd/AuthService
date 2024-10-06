package main

import (
	"authService"
	"github.com/sirupsen/logrus"
)

func main() {
	srv := new(authService.Server)
	if err := srv.Run("8080"); err != nil {
		logrus.Fatalf("error %s", err.Error())
	}
}
