package server

import (
	"errors"

	"github.com/rusik69/kv2/pkg/client/client"
	"github.com/sirupsen/logrus"
)

func (c *Command) parseCmd(body []byte) error {
	cmd, key, value := client.Parse(body)
	logrus.Println("cmd:", cmd)
	logrus.Println("cmd length:", len(cmd))
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
		return errors.New("unknown command: " + cmd)
	}
	return nil
}
