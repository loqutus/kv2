package server

// setkey is a function that sets a value by key in the server.
func (s *Server) SetKey(key string, value int64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.kv[key] = value
	return nil
}
