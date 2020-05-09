package main

import (
	"log"
	"github.com/Anarrr/gomicrodev/webservice"
)

func main() {
	if err := webservice.Serve(); err != nil {
		log.Fatal("Server is down. Err:", err)
	}
}