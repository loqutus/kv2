package server

import (
	"errors"

	"github.com/rusik69/kv2/pkg/client/client"
)

func (c *Command) parseCmd(body []byte) error {
	//bodyString = strings.TrimSuffix(bodyString, "\n")
	//bodyString = strings.TrimPrefix(bodyString, " ")
	cmd, key, value := client.Parse(body)
	switch cmd {
	case "set":
		if key == "" || value == "" {
			return errors.New("not enough arguments for set")
		}
		c.Cmd = "set"
		c.Key = key
		c.Value = value
	case "get":
		if key == "" {
			return errors.New("not enough arguments for get")
		}
		c.Cmd = "get"
		c.Key = key
	case "info":
		c.Cmd = "info"
	default:
		return errors.New("unknown command")
	}
	return nil
}
