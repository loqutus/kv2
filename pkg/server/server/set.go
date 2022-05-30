package server

func (s *Server) Set(c Command) error {
	s.kv[c.Key] = c.Value
	return nil
}
