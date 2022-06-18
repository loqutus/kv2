package server

import (
	"github.com/sirupsen/logrus"
)

// SetInNodes is a function that sets a key-pair to nodes.
func (s *Server) SetInNodes(c Command) error {
	nodes := s.chooseNodes()
	for k := range nodes {
		client := s.nodesClients[k]
		err := client.Set(c.Key, c.Value)
		if err != nil {
			logrus.Errorln(err)
			continue
		}
	}
	return nil
}
