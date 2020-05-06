package main

import (
	"context"
	"fmt"
	pb "github.com/Anarr/gomicrodev/proto/auth"
	"github.com/google/uuid"
	"github.com/micro/go-micro"
	"github.com/pkg/errors"
	"log"
)

const ErrValidation = "email and password are required"

type Auth struct{}

func (a *Auth) Login(ctx context.Context, req *pb.AuthRequest, res *pb.AuthResponse) error {

	if req.Email == "" || req.Password == "" {
		return errors.New(ErrValidation)
	}

	res.IsLoggedIn = true
	res.SessionId = uuid.New().String()
	return nil
}

//setup auth client
func runAuthClient(service micro.Service) {
	client := pb.NewAuthServiceClient("auth", service.Client())

	req := &pb.AuthRequest{
		Email:    "anar.rzayev94@gmail.com",
		Password: "123456789",
	}
	res, err := client.Login(context.Background(), req)

	if err != nil {
		log.Fatal("can not login user")
	}

	fmt.Println("User Logged in sucessfully", res)
}

func main() {
	service := micro.NewService(
		micro.Name("auth"),
	)

	//run client
	//fmt.Println("Try request auth server")
	//runAuthClient(service)
	//return

	pb.RegisterAuthServiceHandler(service.Server(), new(Auth))

	if err := service.Run(); err != nil {
		log.Fatal("Error occurs while running auth server", err)
	}
}
