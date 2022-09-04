package server

import (
	"errors"

	"github.com/rusik69/kv2/pkg/server/persistent"
)

// Delete is a function that deletes a key-value pair from the server.
func (s *Server) Del(c Command) error {
	if v, ok := s.kv[c.Key]; ok {
		err := persistent.PersistentInstance.Del(v)
		if err != nil {
			return err
		}
		delete(s.kv, c.Key)
	} else {
		return errors.New("key not found")
	}
}
