package client

import (
	"fmt"
	"net"
)

// Connect connects to the server
func (c *Client) Connect() error {
	fmt.Println("Connecting to server", c.Host+":"+c.Port)
	var err error
	c.Conn, err = net.Dial("tcp", c.Host+":"+c.Port)
	return err
}
