package server

// Set is a function that sets a key-value pair in the server.
func (s *Server) Set(c Command) error {
	s.kv[c.Key] = c.Value
	return nil
}
