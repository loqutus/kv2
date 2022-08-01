package client

import (
	"net"
	"time"
)

// Client is a client struct.
type Client struct {
	Host      string
	Port      string
	FilesPort string
	Conn      net.Conn
	Timeout   time.Duration
}
