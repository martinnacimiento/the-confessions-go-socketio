package main

import "./controllers"

func main() {
	server := NewServer(":8000")
	server.handle("/confessions", "GET", controllers.IndexConfession)
	server.handle("/confessions", "POST", controllers.StoreConfession)
	server.listen()
}
