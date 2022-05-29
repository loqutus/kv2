package client

import "errors"

func (c *Client) Set(key, value string) error {
	_, err := c.Conn.Write([]byte("SET " + key + " " + value + "\n"))
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
