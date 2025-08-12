package main

import (
	"log"
	"os"

	"github.com/amorindev/go-tmpl/cmd/api/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	hsp := os.Getenv("HTTP_SERVER_PORT")
	if hsp == "" {
		log.Fatal("environment variable HTTP_SERVER_PORT is not set ")
	}

	httpSrv := server.NewHttpServer(hsp)
	httpSrv.Start()
}
