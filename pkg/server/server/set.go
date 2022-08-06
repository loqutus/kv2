package server

// Set is a function that sets a key-value pair in the server.
func (s *Server) Set(c Command) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.kv[c.Key] = c.Value
	return nil
}
