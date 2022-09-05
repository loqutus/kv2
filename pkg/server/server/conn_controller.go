package server

import (
	"time"

	"github.com/rusik69/kv2/pkg/client/client"
)

// ConnController is the main connections controller.
func (s *Server) ConnController() {
	s.nodes = make(map[string]client.Client)
	for _, n := range s.nodeNames {
		var cli client.Client
		cli.Host = n[0]
		cli.Port = n[1]
		err := cli.Connect()
		if err != nil {
			continue
		}
		id, err := cli.GetID()
		if err != nil {
			continue
		}
		if id != s.id {
			s.nodes[id] = cli
		}
	}
	for {
		for id, cli := range s.nodes {
			err := cli.Ping()
			if err != nil {
				delete(s.nodes, id)
				err := cli.Connect()
				if err != nil {
					continue
				}
				id, err := cli.GetID()
				if err != nil {
					continue
				}
				s.nodes[id] = cli
			}
		}
		time.Sleep(time.Second * 5)
	}
}
