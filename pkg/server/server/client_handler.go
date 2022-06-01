package server

import (
	"net"

	"github.com/sirupsen/logrus"
)

// ServerHandler is the main server handler.
func (s *Server) ClientHandler(conn net.Conn) {
	logrus.Println("client handler")
	for {
		respBody := make([]byte, 1024)
		_, err := conn.Read(respBody)
		if err != nil {
			logrus.Errorln(err)
			break
		}
		logrus.Println(string(respBody))
		var c Command
		err = c.parseCmd(respBody)
		if err != nil {
			logrus.Errorln(err)
			conn.Write([]byte(err.Error()))
			continue
		}
		switch c.Cmd {
		case "set":
			err = s.Set(c)
			if err != nil {
				logrus.Errorln(err)
				conn.Write([]byte(err.Error()))
				continue
			} else {
				logrus.Println("SET OK")
				conn.Write([]byte("OK"))
				continue
			}
		case "get":
			value, err := s.Get(c)
			if err != nil {
				logrus.Errorln(err)
				conn.Write([]byte(err.Error()))
				continue
			} else {
				logrus.Println("GET OK")
				conn.Write([]byte(value))
				continue
			}
		}
	}
	return
}
