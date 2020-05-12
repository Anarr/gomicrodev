package main

import (
	"context"
	u "github.com/Anarr/gomicrodev/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserService struct{}

var users []*u.User

func(us *UserService) All(ctx context.Context, req *u.AllRequest) (*u.AllResponse, error) {
	return nil, nil
}

func (us *UserService) Get(ctx context.Context, req *u.GetUserRequest) (*u.User, error) {
	return nil, nil
}

func (us *UserService) Upload(stream u.UserService_UploadServer) error {
	return nil
}

func init() {
	u := []*u.User {
		*u.User{
			Id:1,
			Name: "Anar Rzayev",
			Username: "anar.rzayev",
		},
		*u.User{
			Id:2,
			Name:"Elmir Mahmudov",
			Username:"elmirm",
		},
		*u.User{
			Id:3,
			Name: "Elshad Yarmetov",
			Username:"elsh",
		},
	}

	user = append(users)
}

func main() {
	list, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	u.RegisterUserServiceServer(server, new(UserService))

	if err = server.Serve(list); err != nil {
		log.Fatal(err)
	}
}
