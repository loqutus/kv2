package client

import "github.com/rusik69/kv2/pkg/client/argparse"

// Init initializes a client.
func (c *Client) Init(arguments argparse.Args) {
	c.Host = arguments.ServerHost
	c.Port = arguments.ServerPort
}
