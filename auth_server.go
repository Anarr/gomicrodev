package main

import (
	"context"
	"fmt"
	pb "github.com/Anarr/gomicrodev/proto/auth"
	"github.com/google/uuid"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/client/http"
	"github.com/pkg/errors"
	"log"
)

const ErrValidation = "email and password are required"

type Auth struct{}

func (a *Auth) Login(ctx context.Context, req *pb.AuthRequest, res *pb.AuthResponse) error {

	if err := a.Validate(ctx, req, &pb.ValidationResponse{}); err != nil {
		return err
	}

	res.IsLoggedIn = true
	res.SessionId = uuid.New().String()
	return nil
}

func (a *Auth) Validate(ctx context.Context, req *pb.AuthRequest, res *pb.ValidationResponse) error {
	if req.Email == "" || req.Password == "" {
		res.IsOk = false
		log.Println("validation err", res)
		return errors.New(ErrValidation)
	}

	res.IsOk = true
	return nil
}

//setup auth client
func runAuthClient(service micro.Service) {
	client := pb.NewAuthServiceClient("auth", service.Client())

	req := &pb.AuthRequest{
		Email:    "anar.rzayev94@gmail.com",
		Password: "124",
	}
	res, err := client.Login(context.Background(), req)

	if err != nil {
		log.Println("login responses", res)
		log.Fatal("login action finished with err: ", err)

	}

	fmt.Println("User Logged in sucessfully", res)
}

func main() {
	service := micro.NewService(
		micro.Name("auth"),
		micro.Client(http.NewClient()),
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
