package client

import (
	"time"

	"github.com/rusik69/kv2/pkg/client/argparse"
)

// Init initializes a client.
func (c *Client) Init(arguments argparse.Args) {
	c.Host = arguments.ServerHost
	c.Port = arguments.ServerPort
	c.FilesPort = arguments.ServerFilesPort
	c.Timeout = time.Second * 60
}
