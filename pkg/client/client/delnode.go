package client

import "errors"

// DelNode is a function that adds a node to the nodes list.
func (c *Client) DelNode(host, port string) error {
	connBody := []byte("delnode " + host + " " + port)
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
