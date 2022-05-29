package client

import "net"

type Client struct {
	Host string
	Port string
	Conn net.Conn
}

var ClientInstance Client
