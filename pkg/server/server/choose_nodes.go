package server

import "github.com/rusik69/kv2/pkg/client/client"

// choose nodes to set a key-pair to
func (s *Server) chooseNodes() map[string]client.Client {
	if s.replicas < len(s.nodes) {
		return s.nodes
	}
	nodes := make(map[string]client.Client, s.replicas)
	for i := 0; i < s.replicas-1; i++ {
		for k, v := range s.nodes {
			nodes[k] = v
			break
		}
	}
	return nodes
}
