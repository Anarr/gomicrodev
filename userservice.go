package main

import (
	"log"
	"net"
	u "github.com/Anarr/gomicrodev/proto/user"
	"google.golang.org/grpc"
)

type UserService struct{}

func main() {
	list, err := net.Listen("tcp",":90909")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	u.RegisterUserServiceHandler(server, new(UserService))

	if err = server.Serve(list); err != nil {
		log.Fatal(err)
	}
}
