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
		if c.Key == "" || c.Value == nil {
			return errors.New("not enough arguments for set")
		}
	case "get":
		if c.Key == "" {
			return errors.New("not enough arguments for get")
		}
	case "info":
		break
	case "del":
		if c.Key == "" {
			return errors.New("not enough arguments for del")
		}
	case "addnode":
		if c.Key == "" {
			return errors.New("not enough arguments for addnode")
		}
	case "delnode":
		if c.Key == "" {
			return errors.New("not enough arguments for delnode")
		}
	case "setreplicas":
		if c.Key == "" {
			return errors.New("not enough aruments for setreplicas")
		}
	case "ping":
		break
	case "id":
		break
	default:
		return errors.New("unknown command: " + c.Cmd)
	}
	return nil
}
