package client

import (
	"errors"
	"net/http"
	"os"
)

// Download is a function that downloads a file from the server.
func (c *Client) Download(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	url := "http://" + c.Host + ":" + c.FilesPort + "/" + fileName
	client := http.Client{
		Timeout: c.Timeout,
	}
	req, err := http.NewRequest("GET", url, f)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}
