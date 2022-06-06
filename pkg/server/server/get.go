package server

import (
	"errors"
)

// Get is a function that gets a value by key from the server.
func (s *Server) Get(c Command) (string, error) {
	//logrus.Println("key length:", len(c.Key))
	//logrus.Println("key:", c.Key)
	if v, ok := s.kv[c.Key]; ok {
		return v, nil
	} else {
		return "", errors.New("key not found")
	}
}
