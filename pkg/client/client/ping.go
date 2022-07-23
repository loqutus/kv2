package client

import "errors"

// Ping sends a ping to the server and returns the response.
func (c Client) Ping() error {
	connBody := []byte("ping")
	_, err := c.Conn.Write(connBody)
	if err != nil {
		return err
	}
	respBody := make([]byte, 4)
	_, err = c.Conn.Read(respBody)
	if err != nil {
		return err
	}
	if string(respBody) != "pong" {
		return errors.New("ping failed, got " + string(respBody))
	}
	return nil
}
