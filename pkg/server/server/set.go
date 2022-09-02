package server

import "github.com/rusik69/kv2/pkg/server/persistent"

// Set is a function that sets a key-value pair in the server.
func (s *Server) Set(c Command) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	data := persistent.BinaryData{Key: []byte(c.Key), Val: c.Value, Del: false}
	offset, err := persistent.PersistentInstance.Set(data)
	if err != nil {
		return err
	}
	s.kv[c.Key] = offset
	return nil
}
