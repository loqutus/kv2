package client

//GetID returns the server id
func (c *Client) GetID() (string, error) {
	connBody := []byte("id")
	_, err := c.Conn.Write(connBody)
	if err != nil {
		return "", err
	}
	respBody := make([]byte, 36)
	_, err = c.Conn.Read(respBody)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}
