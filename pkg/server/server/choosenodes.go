package server

// choose nodes to set a key-pair to
func (s *Server) chooseNodes() map[string]Node {
	if s.replicas < len(s.nodes) {
		return s.nodes
	}
	nodes := make(map[string]Node, s.replicas)
	for i := 0; i < s.replicas-1; i++ {
		for k, v := range s.nodes {
			nodes[k] = v
			break
		}
	}
	return nodes
}
