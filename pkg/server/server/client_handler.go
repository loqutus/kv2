package server

import (
	"io"
	"net"

	"github.com/sirupsen/logrus"
)

// ServerHandler is the main server handler.
func (s *Server) ClientHandler(conn net.Conn) {
	logrus.Println("client handler")
	for {
		c, err := s.GetCmd(conn)
		if err == io.EOF {
			continue
		} else if err != nil {
			break
		}
		switch c.Cmd {
		case "set":
			err = s.Set(c)
			if err != nil {
				logrus.Errorln(err)
				conn.Write([]byte(err.Error()))
				continue
			} else {
				if len(s.nodes) > 0 {
					err = s.SetInNodes(c)
					if err != nil {
						logrus.Errorln(err)
						conn.Write([]byte(err.Error()))
						continue
					}
				}
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
				conn.Write(value)
				continue
			}
		case "del":
			err = s.Del(c)
			if err != nil {
				logrus.Errorln(err)
				conn.Write([]byte(err.Error()))
				continue
			} else {
				conn.Write([]byte("OK"))
				continue
			}
		case "info":
			info, err := s.Info()
			if err != nil {
				logrus.Errorln(err)
				conn.Write([]byte(err.Error()))
				continue
			} else {
				conn.Write([]byte(info))
				continue
			}
		case "addnode":
			err = s.AddNode(c, true)
			if err != nil {
				logrus.Errorln(err)
				conn.Write([]byte(err.Error()))
				continue
			} else {
				conn.Write([]byte("OK"))
				continue
			}
		case "delnode":
			err = s.DelNode(c)
			if err != nil {
				logrus.Errorln(err)
				conn.Write([]byte(err.Error()))
				continue
			} else {
				conn.Write([]byte("OK"))
				continue
			}
		case "setreplicas":
			err = s.SetReplicas(c)
			if err != nil {
				logrus.Errorln(err)
				conn.Write([]byte(err.Error()))
				continue
			} else {
				conn.Write([]byte("OK"))
				continue
			}
		case "id":
			conn.Write([]byte(s.id))
		default:
			logrus.Errorln("unknown command:", c.Cmd)
			conn.Write([]byte("unknown command: " + c.Cmd))
		}
	}
}
