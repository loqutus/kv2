package client

func (c *Client) Get(key string) ([]byte, error) {
	_, err := c.Conn.Write([]byte("get " + key))
	if err != nil {
		return nil, err
	}
	respBody := make([]byte, 1024)
	_, err = c.Conn.Read(respBody)
	if err != nil {
		return nil, err
	}
	response := []byte{}
	for i := 0; i < len(respBody); i++ {
		if respBody[i] == 0 {
			break
		}
		response = append(response, respBody[i])
	}
	return response, nil
}
