package webservice

import (
	"context"
	"encoding/json"
	"fmt"
	au "github.com/Anarr/gomicrodev/proto/auth"
	gr "github.com/Anarr/gomicrodev/proto/greeter"
	ps "github.com/Anarr/gomicrodev/proto/post"
	"github.com/micro/go-micro"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	authService, greetingService, postService micro.Service
	asc au.AuthServiceClient
	gsc gr.GreeterClient
	psc ps.PostServiceClient
	//userClient u.UserServiceClient
)

var listenAddr = ":8080"

type User struct {
	ID   int
	Name string
}

//init initalize servicec
func init() {
	authService = micro.NewService(micro.Name("auth"))
	greetingService = micro.NewService(micro.Name("greeter"))
	postService = micro.NewService(micro.Name("post"))

}

//init initalizes service clients
func init()  {
	asc = au.NewAuthServiceClient("auth", authService.Client())
	gsc = gr.NewGreeterClient("greeter", greetingService.Client())
	psc = ps.NewPostServiceClient("post", postService.Client())
}

func loginHandler(writer http.ResponseWriter, req *http.Request) {
	email := req.URL.Query().Get("email")
	pass := req.URL.Query().Get("pass")

	r := &au.AuthRequest{
		Email:    email,
		Password: pass,
	}
	res, err := asc.Login(context.Background(), r)
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
	res, err := gsc.Hello(context.Background(), &gr.Request{Name: "Anar"})
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
		r := &ps.PostsRequest{}
		res, err := psc.All(context.Background(), r)

		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		if err != nil {
			fmt.Fprint(writer, err)
			return
		}

		posts, _ := json.Marshal(&res)
		writer.Write(posts)
	case "post":
		r := &ps.PostCreateRequest{
			UserId:      1,
			Description: req.URL.Query().Get("description"),
			CreatedAt:   time.Now().String(),
		}

		res, err := psc.Create(context.Background(), r)

		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)

		if err != nil {
			fmt.Fprint(writer, err)
			return
		}

		post, _ := json.Marshal(res)

		writer.Write(post)

	case "delete":
		postId, _ := strconv.Atoi(req.URL.Query().Get("post_id"))
		r := &ps.PostDeleteRequest{
			UserId:      1,
			PostId: int64(postId),
		}

		res, err := psc.Delete(context.Background(), r)

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

func usersHandler(writer http.ResponseWriter, _ *http.Request) {

}

func Serve() error {

	fmt.Println("Starting web server...")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("microservices"))
	})
	//auth handlers
	http.HandleFunc("/auth/login", loginHandler)

	//greeting handlers
	http.HandleFunc("/greeting", greetingHandler)

	//posts handlers
	http.HandleFunc("/posts", postsHandler)

	//users handlers
	http.HandleFunc("/users", usersHandler)

	fmt.Println("Server is ready.")
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		return err
	}

	return nil
}
