package main

import (
	"github.com/Anarr/gomicrodev/postservice"
	"log"
)
func main() {
	if err := postservice.Serve(); err != nil {
		log.Fatal("Register post server error:", err)
	}
}