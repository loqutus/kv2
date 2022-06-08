package client

// Info returns information about the server.
func (c *Client) Info() (string, error) {
	_, err := c.Conn.Write([]byte("info"))
	if err != nil {
		return "", err
	}
	respBody := make([]byte, 1024)
	_, err = c.Conn.Read(respBody)
	if err != nil {
		return "", err
	}
	var response string
	for i := 0; i < len(respBody); i++ {
		if respBody[i] == 0 {
			break
		}
		response += string(respBody[i])
	}
	return response, nil
}
