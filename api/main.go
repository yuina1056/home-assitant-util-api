package main

import (
	"home-assitant-util-api/api/server"
	"os"
)

func main() {
	server.Handler(os.Getenv("PORT"))
}
