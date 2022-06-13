package server

import (
	"net"
	"sync"
)

// Server is the main server struct.
type Server struct {
	listenPortClient string
	listenPortServer string
	listenHost       string
	listenerServer   net.Listener
	listenerClient   net.Listener
	wg               sync.WaitGroup
	kv               map[string][]byte
}

var ServerInstance Server

// Command struct
type Command struct {
	Cmd   string
	Key   string
	Value []byte
}
