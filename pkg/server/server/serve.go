package server

import (
	"log"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Serve starts the server.
func (s *Server) Serve() {
	logrus.Info("starting server")
	if s.debug {
		startPyroscope()
	}
	s.kv = make(map[string][]byte)
	s.id = uuid.New().String()
	s.replicas = 1
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
		}
		logrus.Println("server connected")
		go s.ServerHandler(conn)
	}
	s.wg.Done()
}

// ServeClient starts the client connection handler loop.
func (s *Server) ServeClient() {
	defer s.listenerClient.Close()
	for {
		logrus.Println("waiting for client")
		conn, err := s.listenerClient.Accept()
		if err != nil {
			log.Fatal(err)
		}
		logrus.Println("client connected")
		go s.ClientHandler(conn)
	}
	s.wg.Done()
}
