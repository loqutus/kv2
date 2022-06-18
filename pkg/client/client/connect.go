package client

import (
	"net"
)

// Connect connects to the server
func (c *Client) Connect() error {
	var err error
	c.Conn, err = net.Dial("tcp", c.Host+":"+c.Port)
	return err
}
