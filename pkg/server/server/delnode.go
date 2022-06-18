package server

import "errors"

func (s *Server) DelNode(c Command) error {
	host := c.Key
	port := string(c.Value)
	deleted := false
	for id, v := range s.nodes {
		if v.Host == host && v.Port == port {
			delete(s.nodes, id)
			delete(s.nodesClients, id)
			deleted = true
			break
		}
	}
	if deleted {
		return nil
	} else {
		return errors.New("node not found")
	}
}
