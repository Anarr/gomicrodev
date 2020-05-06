package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	"net/http"
	au "github.com/Anarr/gomicrodev/proto/auth"
	gr "github.com/Anarr/gomicrodev/proto/greeter"
)

var authService, greetingService micro.Service

var listenAddr = ":8080"
type User struct {
	ID int
	Name string
}

func init () {
	authService = micro.NewService(micro.Name("auth"))
	greetingService = micro.NewService(micro.Name("greeter"))
}

func loginHandler(writer http.ResponseWriter, req *http.Request) {
	email := req.URL.Query().Get("email")
	pass := req.URL.Query().Get("pass")

	client := au.NewAuthServiceClient("auth", authService.Client())
	r := &au.AuthRequest{
		Email: email,
		Password:pass,
	}
	res, err := client.Login(context.Background(), r)
	writer.Header().Add("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err != nil {
		fmt.Fprint(writer, err)
		return
	}

	js, _ := json.Marshal(&res)
	writer.Write(js)
}
func greetingHandler(writer http.ResponseWriter, req *http.Request) {
	client := gr.NewGreeterClient("greeter", greetingService.Client())
	res, err := client.Hello(context.Background(), &gr.Request{Name: "Anar"})
	writer.Header().Add("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if err != nil {
		fmt.Fprint(writer, err)
		return
	}

	greeting, _ := json.Marshal(res)
	writer.Write(greeting)
	return
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/greeting", greetingHandler)
	if err := http.ListenAndServe(listenAddr, nil); err !=nil {
		log.Fatal("Web server down")
	}
}
