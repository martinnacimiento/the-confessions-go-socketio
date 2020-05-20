package main

import (
	"log"
	"net/http"
)

// NewServer create a new server
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// Server a server
type Server struct {
	port   string
	router *Router
}

func (s *Server) listen() error {
	http.Handle("/", s.router) // Agrega todas las rutas
	//NewWebsocket()
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		log.Fatal("unable to run server", err)
		return err
	}
	log.Print("server running on", s.port)
	return nil
}

func (s *Server) handle(path string, method string, handler http.HandlerFunc) {
	//Check if the path already exists
	if !s.router.findPath(path) {
		//If not path then create a new one
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s *Server) addMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Handler handler
type Handler func(w http.ResponseWriter, r *http.Request)

// Middleware middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc
