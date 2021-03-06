package main

import (
	"context"
	pb "github.com/Anarr/gomicrodev/proto/greeter"
	"github.com/micro/go-micro"
	"log"
)

//Define greeter service methods
type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, res *pb.Response) error {
	res.Greeting = "Hello," + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"for": "test reason",
		}),
	)

	// Setup the server
	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Println("Greeter server error:", err)
	}
}
