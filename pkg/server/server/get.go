package server

import (
	"errors"
)

// Get is a function that gets a value by key from the server.
func (s *Server) Get(c Command) ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if v, ok := s.kv[c.Key]; ok {
		return s.kv, nil
	} else {
		return nil, errors.New("key not found")
	}
}
