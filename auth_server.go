package main

import (
	"github.com/Anarr/gomicrodev/authservice"
	"log"
)
func main() {
	if err := authservice.Serve(); err != nil {
		log.Fatal("Register auth server error:", err)
	}
}