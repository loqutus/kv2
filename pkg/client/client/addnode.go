package client

import "errors"

// AddNode is a function that adds a node to the nodes list.
func (c *Client) AddNode(node, port string) error {
	connBody := []byte("addnode " + node + " " + port)
	_, err := c.Conn.Write(connBody)
	if err != nil {
		return err
	}
	respBody := make([]byte, 2)
	_, err = c.Conn.Read(respBody)
	if err != nil {
		return err
	}
	if string(respBody) != "OK" {
		return errors.New("server error : " + err.Error())
	}
	return nil
}
