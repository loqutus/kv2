package server

import "github.com/rusik69/kv2/pkg/client/client"

// AddNode is a function that adds a node to the nodes list.
func (s *Server) AddNode(c Command, fromClient bool) error {
	host := c.Key
	port := string(c.Value)
	var cli client.Client
	cli.Host = host
	cli.Port = port
	cli.Connect()
	id, err := cli.GetID()
	if err != nil {
		return err
	}
	if fromClient {
		err = cli.AddNode(s.listenHost, s.listenPortServer)
		if err != nil {
			return err
		}
	}
	s.nodes[id] = cli
	return nil
}
