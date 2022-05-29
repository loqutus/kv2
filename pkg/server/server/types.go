package server

import "net"

// Server is the main server struct.
type Server struct {
	listenPortClient string
	listenPortServer string
	listenHost       string
	listenerServer   net.Listener
	listenerClient   net.Listener
}

var ServerInstance Server
