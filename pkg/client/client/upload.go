package client

import (
	"errors"
	"net/http"
	"os"
)

// Upload is a function that uploads a file to the server.
func (c *Client) Upload(source string, destination string) error {
	f, err := os.Open(source)
	if err != nil {
		return err
	}
	defer f.Close()
	url := "http://" + c.Host + ":" + c.FilesPort + "/" + destination
	client := http.Client{
		Timeout: c.Timeout,
	}
	req, err := http.NewRequest("POST", url, f)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}
