package server

import (
	"io"
	"net"

	"github.com/sirupsen/logrus"
)

// GetCmd reads and parses command from client connection.
func (s *Server) GetCmd(conn net.Conn) (Command, error) {
	var c Command
	respBody := make([]byte, 1024)
	_, err := conn.Read(respBody)
	if err == io.EOF {
		return c, err
	}
	if err != nil {
		logrus.Errorln(err)
		return c, err
	}
	err = c.parseCmd(respBody)
	if err != nil {
		logrus.Errorln(err)
		conn.Write([]byte(err.Error()))
		return c, nil
	}
	return c, nil
}
