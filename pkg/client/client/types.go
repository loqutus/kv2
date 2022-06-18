package client

import "net"

// Client is a client struct.
type Client struct {
	Host string
	Port string
	Conn net.Conn
}
