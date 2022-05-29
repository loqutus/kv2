package server

import (
	"net"
	"os"

	"github.com/sirupsen/logrus"
)

// Listen returns listeners for client/server connections.
func (s *Server) Listen() (listenerClient net.Listener, listenerServer net.Listener) {
	listenerClient, err := net.Listen("tcp4", s.listenHost+":"+s.listenPortClient)
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
	listenerServer, err = net.Listen("tcp4", s.listenHost+":"+s.listenPortServer)
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
	return
}
