package client

import "errors"

// SetReplicas is a function that sets the number of replicas in the server.
func (c *Client) SetReplicas(replicas string) error {
	connBody := []byte("setreplicas " + replicas)
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
