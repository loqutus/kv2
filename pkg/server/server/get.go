package server

import (
	"errors"

	"github.com/rusik69/kv2/pkg/server/persistent"
)

// Get is a function that gets a value by key from the server.
func (s *Server) Get(c Command) ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if offset, ok := s.kv[c.Key]; ok {
		return persistent.PersistentInstance.Get(offset)
	} else {
		return nil, errors.New("key not found")
	}
}
