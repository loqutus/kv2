package server

// AddNode is a function that adds a node to the nodes list.
func (s *Server) AddNode(c Command) error {
	host := c.Key
	port := string(c.Value)
	s.nodes[host] = port
	return nil
}
