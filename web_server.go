package main

import (
	"context"
	"encoding/json"
	"fmt"
	au "github.com/Anarr/gomicrodev/proto/auth"
	gr "github.com/Anarr/gomicrodev/proto/greeter"
	ps "github.com/Anarr/gomicrodev/proto/post"
	"github.com/micro/go-micro"
	"log"
	"net/http"
	"strings"
	"time"
)

var authService, greetingService, postService micro.Service

var listenAddr = ":8080"

type User struct {
	ID   int
	Name string
}

func init() {
	authService = micro.NewService(micro.Name("auth"))
	greetingService = micro.NewService(micro.Name("greeter"))
	postService = micro.NewService(micro.Name("post"))
}

func loginHandler(writer http.ResponseWriter, req *http.Request) {
	email := req.URL.Query().Get("email")
	pass := req.URL.Query().Get("pass")

	client := au.NewAuthServiceClient("auth", authService.Client())
	r := &au.AuthRequest{
		Email:    email,
		Password: pass,
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
func greetingHandler(writer http.ResponseWriter, _ *http.Request) {
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

func postsHandler(writer http.ResponseWriter, req *http.Request) {
	method := strings.ToLower(req.Method)

	switch method {
	case "get":
		client := ps.NewPostServiceClient("post", postService.Client())
		r := &ps.PostsRequest{}
		res, err := client.All(context.Background(), r)

		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		if err != nil {
			fmt.Fprint(writer, err)
			return
		}

		posts, _ := json.Marshal(&res)
		writer.Write(posts)
	case "post":
		client := ps.NewPostServiceClient("post", postService.Client())

		r := &ps.PostCreateRequest{
			UserId:      1,
			Description: req.URL.Query().Get("desc"),
			CreatedAt:   time.Now().String(),
		}

		res, err := client.Create(context.Background(), r)

		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)

		if err != nil {
			fmt.Fprint(writer, err)
			return
		}

		post, _ := json.Marshal(res)

		writer.Write(post)
		return
	default:
		writer.Write([]byte("method not allowed by app"))
	}

	return
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("microservices"))
	})
	//auth handlers
	http.HandleFunc("/auth/login", loginHandler)
	//greeting handlers
	http.HandleFunc("/greeting", greetingHandler)
	//posts handlers
	http.HandleFunc("/posts", postsHandler)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal("Web server down")
	}
}
