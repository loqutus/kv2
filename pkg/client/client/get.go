package client

func (c *Client) Get(key string) (string, error) {
	_, err := c.Conn.Write([]byte("get " + key))
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
