package server

import (
	"log"
	"net/http"
)

// Server is a handler around http.Server
type Server struct {
	ApiServer *http.Server
}

func NewServer(handler http.Handler) *Server {
	return &Server{
		ApiServer: &http.Server{
			Handler: handler,
		},
	}
}

// Serve method will start our api-server for handler on the provided address
func (s *Server) Serve(addr string) {
	s.ApiServer.Addr = addr
	// TODO: start on a new routine
	log.Printf("Starting Server on %s", addr)
	s.listenServer()
}

func (s *Server) listenServer() {
	err := s.ApiServer.ListenAndServe()
	if err != nil {
		// TODO: log fatal
	}
}

func (s *Server) WaitForShutdown() {
	// TODO::
}
