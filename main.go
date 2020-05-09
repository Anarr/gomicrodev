package main

import (
	"github.com/Anarr/go-micro/greeterservice"
	"log"
)

func main() {
	if err := greeterservice.Serve(); err != nil {
		log.Fatal("Register greeting server error:", err)
	}
}