package main

import (
	"home-assitant-util-api/api/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.Handler(os.Getenv("PORT"))
}
