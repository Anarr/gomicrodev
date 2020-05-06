package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	"net/http"
	pb "github.com/Anarr/gomicrodev/proto/auth"
)
var listenAddr = ":8080"
type User struct {
	ID int
	Name string
}
func loginHandler(writer http.ResponseWriter, req *http.Request) {
	var s micro.Service
	email := req.URL.Query().Get("email")
	pass := req.URL.Query().Get("pass")

	client := pb.NewAuthServiceClient("auth", s.Client())
	r := &bp.AuthRequest{
		Email: email,
		Password:pass,
	}
	res, err := client.Login(context.Background(), r)

	if err != nil {
		fmt.Fprint(writer, err)
	}

	writer.Header().Add("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	js, _ := json.Marshal(&res)
	writer.Write(js)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	if err := http.ListenAndServe(listenAddr, nil); err !=nil {
		log.Fatal("Web server down")
	}
}