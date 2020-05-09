package main

import (
	"github.com/micro/go-micro"
	u "github.com/Anarr/gomicrodev/proto/user"
	"log"
)

type UserService struct {}

func main() {
	service := micro.NewService(
		micro.Name("user"),
	)

	u.NewUserServiceHandler(service.Server(), new(UserService))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}