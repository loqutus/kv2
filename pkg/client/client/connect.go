package client

import (
	"net"
	"os"

	"github.com/sirupsen/logrus"
)

func (c *Client) Connect() {
	var err error
	c.Conn, err = net.Dial("tcp", c.Host+":"+c.Port)
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
