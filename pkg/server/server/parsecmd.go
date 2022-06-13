package server

import (
	"errors"

	"github.com/rusik69/kv2/pkg/client/client"
)

// parseCmd is parsing command to Command object
func (c *Command) parseCmd(body []byte) error {
	c.Cmd, c.Key, c.Value = client.Parse(body)
	switch c.Cmd {
	case "set":
		if c.Key == "" || len(c.Value) == 0 {
			return errors.New("not enough arguments for set")
		}
	case "get":
		if c.Key == "" {
			return errors.New("not enough arguments for get")
		}
	case "info":
		return nil
	case "del":
		if c.Key == "" {
			return errors.New("not enough arguments for delete")
		}
	default:
		return errors.New("unknown command: " + c.Cmd)
	}
	return nil
}
