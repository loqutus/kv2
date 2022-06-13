package server

import "strconv"

// SetReplicas is a function that sets replicas number in a server
func (s *Server) SetReplicas(c Command) error {
	replicas, err := strconv.Atoi(c.Key)
	if err != nil {
		return err
	}
	s.replicas = replicas
	return nil
}
