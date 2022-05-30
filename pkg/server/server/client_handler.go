package server

import (
	"net"

	"github.com/sirupsen/logrus"
)

// ServerHandler is the main server handler.
func (s *Server) ClientHandler(conn net.Conn) {
	respBody := make([]byte, 1024)
	_, err := conn.Read(respBody)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	var c Command
	err = c.parseCmd(respBody)
	if err != nil {
		logrus.Errorln(err)
		conn.Write([]byte(err.Error()))
		return
	}
	switch c.Cmd {
	case "SET":
		err = s.Set(c)
		if err != nil {
			logrus.Errorln(err)
			conn.Write([]byte(err.Error()))
			return
		} else {
			conn.Write([]byte("OK"))
			return
		}
	case "GET":
		value, err := s.Get(c)
		if err != nil {
			logrus.Errorln(err)
			conn.Write([]byte(err.Error()))
			return
		}
		conn.Write([]byte(value))
	}
	return
}
