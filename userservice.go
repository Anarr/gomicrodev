package main

import (
	u "github.com/Anarr/gomicrodev/proto/user"
	"github.com/micro/go-micro"
	"log"
)

type UserService struct{}

func main() {
	service := micro.NewService(
		micro.Name("user"),
	)

	u.RegisterUserServiceHandler(service.Server(), new(UserService))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
