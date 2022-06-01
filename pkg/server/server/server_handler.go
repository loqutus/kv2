package server

import (
	"net"

	"github.com/sirupsen/logrus"
)

// ServerHandler is the main server handler.
func (s *Server) ServerHandler(conn net.Conn) {
	logrus.Println("server handler")
	return
}
