package main

import (
	"home-assitant-util-api/api/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") == "local" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
	server.Handler(os.Getenv("PORT"))
}
