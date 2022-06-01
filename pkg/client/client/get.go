package client

func (c *Client) Get(key string) (string, error) {
	_, err := c.Conn.Write([]byte("GET " + key))
	if err != nil {
		return "", err
	}
	respBody := make([]byte, 1024)
	_, err = c.Conn.Read(respBody)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}
