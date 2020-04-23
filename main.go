package main

import (
	"os"

	"./controllers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}
	server := NewServer(":" + port)
	server.handle("/confessions", "GET", controllers.IndexConfession)
	server.handle("/confessions", "POST", controllers.StoreConfession)
	server.listen()
}
