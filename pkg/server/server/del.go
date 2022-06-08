package server

import (
	"errors"
)

// Delete is a function that deletes a key-value pair from the server.
func (s *Server) Del(c Command) error {
	if _, ok := s.kv[c.Key]; ok {
		delete(s.kv, c.Key)
		return nil
	} else {
		return errors.New("key not found")
	}
}
