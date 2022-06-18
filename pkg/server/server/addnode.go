package server

import "github.com/rusik69/kv2/pkg/client/client"

// AddNode is a function that adds a node to the nodes list.
func (s *Server) AddNode(c Command) error {
	host := c.Key
	port := string(c.Value)
	node := Node{Host: host, Port: port}
	var cli client.Client
	cli.Host = host
	cli.Port = port
	cli.Connect()
	id, err := cli.GetID()
	if err != nil {
		return err
	}
	s.nodes[id] = node
	return nil
}
