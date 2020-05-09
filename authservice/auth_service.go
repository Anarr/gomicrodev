package main

import (
	"context"
	pb "github.com/Anarr/gomicrodev/proto/auth"
	"github.com/google/uuid"
	"github.com/micro/go-micro"
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

func main() {
	service := micro.NewService(
		micro.Name("auth"),
	)

	pb.RegisterAuthServiceHandler(service.Server(), new(Auth))

	if err := service.Run(); err != nil {
		log.Println("Auth server error", err)
	}
}
