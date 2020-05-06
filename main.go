package main

import (
	pb "github.com/Anarr/proto"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"os"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, res *pb.Response) error {
	res.Greeting = "Hello," + req.Name
	return nil
}

//setup client

func runClient(service micro.Service) {
	//create new greeter client
	client := pb.NewGreeterClient("greeter", service.Client())
	res, err := client.Hello(context.Background(), &pb.Request{Name:"Anar"})
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
			"for":"test reason",
		}),
		micro.Flags(&cli.BoolFlag{
			Name: "run_client",
			Usage: "Launch client",
		}),
	)

	service.Init(
		micro.Action(func(c *cli.Context) error {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
			return nil
		}),
	)

	// By default we'll run the server unless the flags catch us

	// Setup the server

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal("Error while running server:", err)
	}
}
