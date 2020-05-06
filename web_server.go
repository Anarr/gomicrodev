package main

import (
	"encoding/json"
	"log"
	"net/http"
)
var listenAddr = ":8080"
type User struct {
	ID int
	Name string
}
func loginHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(&User{ID:12, Name:"Anar Rzayev"})
	writer.Write(res)
	return
}

func main() {
	http.HandleFunc("/login", loginHandler)
	if err := http.ListenAndServe(listenAddr, nil); err !=nil {
		log.Fatal("Web server down")
	}
}