package server

func (s *Server) DelNode(c Command) error {
	host := c.Key
	delete(s.nodes, host)
	return nil
}
