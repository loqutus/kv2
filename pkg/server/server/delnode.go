package server

import "errors"

// DelNode is a function that deletes a node from the server.
func (s *Server) DelNode(c Command) error {
	host := c.Key
	port := string(c.Value)
	deleted := false
	for id, v := range s.nodes {
		if v.Host == host && v.Port == port {
			delete(s.nodes, id)
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
