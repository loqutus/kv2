package server

import "errors"

func (s *Server) Get(c Command) (string, error) {
	if v, ok := s.kv[c.Key]; ok {
		return v, nil
	} else {
		return "", errors.New("key not found")
	}
}
