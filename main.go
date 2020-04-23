package main

import (
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server := NewServer(":" + port)
	server.handle("/confessions", "GET", IndexConfession)
	server.handle("/confessions", "POST", StoreConfession)
	server.listen()
}
