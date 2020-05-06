package main

import (
	"context"
	"fmt"
	pb "github.com/Anarr/gomicrodev/proto/greeter"
	"github.com/micro/go-micro"
	"log"
)

const (
	ValidationErr = "email and password are required"
)

//Define greeter service methods
type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, res *pb.Response) error {
	res.Greeting = "Hello," + req.Name
	return nil
}
//setup client

func runClient(service micro.Service) {
	//create new greeter client
	client := pb.NewGreeterClient("greeter", service.Client())
	res, err := client.Hello(context.Background(), &pb.Request{Name: "Anar"})
	if err != nil {
		log.Fatal("Client err", err)
	}

	fmt.Println("Response is: ", res.GetGreeting())
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"for": "test reason",
		}),
	)
	//fmt.Println("Try request to server")
	//runClient(service)

	// By default we'll run the server unless the flags catch us

	// Setup the server

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal("Error occurs during running greeting server", err)
	}
}
