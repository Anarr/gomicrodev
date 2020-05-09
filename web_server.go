package main

import (
	"log"
	"github.com/Anarr/gomicrodev/webservice"
)

func main() {
	if err := webservice.Serve(); err != nil {
		log.Fatal("Server is down. Err:", err)
	}
}