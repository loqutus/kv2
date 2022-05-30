package server

import (
	"errors"
	"strings"
)

func (c *Command) parseCmd(body []byte) error {
	bodySplit := strings.Split(string(body), " ")
	if len(bodySplit) == 0 {
		return errors.New("empty command")
	}
	cmd := bodySplit[0]
	switch cmd {
	case "SET":
		if len(bodySplit) < 3 {
			return errors.New("not enough arguments for set")
		}
		c.Cmd = "set"
		c.Key = bodySplit[1]
		c.Value = strings.Join(bodySplit[2:], " ")
	case "GET":
		if len(bodySplit) < 2 {
			return errors.New("not enough arguments for get")
		}
		c.Cmd = "get"
		c.Key = bodySplit[1]
	default:
		return errors.New("unknown command")
	}
	return nil
}
