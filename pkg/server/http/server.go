package http

import (
	"log"
	"net/http"
)

//Server represents an http server that handle request
type Server struct {
	handler *Handler
}

//NewServer return an http server with a given handler
func NewServer(h *Handler) *Server {
	return &Server{
		handler: h,
	}
}

//Serve launch the webserver
func (s *Server) Serve() {
	log.Fatal(http.ListenAndServe(":9000", s.handler.Router()))
}
