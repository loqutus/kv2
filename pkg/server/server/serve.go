package server

import (
	"log"
	"os"
)

// Serve starts the server.
func (s *Server) Serve() {
	s.listenerClient, s.listenerServer = s.Listen()
	go s.ServeServer()
	go s.ServeClient()
}

// ServeServer starts the server connection handler loop.
func (s *Server) ServeServer() {
	defer s.listenerServer.Close()
	for {
		conn, err := s.listenerServer.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go s.ServerHandler(conn)
	}
}

// ServeClient starts the client connection handler loop.
func (s *Server) ServeClient() {
	defer s.listenerClient.Close()
	for {
		conn, err := s.listenerClient.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go s.ClientHandler(conn)
	}
}
