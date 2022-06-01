package server

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

// Serve starts the server.
func (s *Server) Serve() {
	logrus.Println("starting server")
	s.kv = make(map[string]string)
	s.listenerClient, s.listenerServer = s.Listen()
	s.wg.Add(2)
	go s.ServeServer()
	go s.ServeClient()
	s.wg.Wait()
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
	s.wg.Done()
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
	s.wg.Done()
}
