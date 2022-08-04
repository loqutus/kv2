package client

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Download is a function that downloads a file from the server.
func (c *Client) Download(fileName string) error {
	url := "http://" + c.Host + ":" + c.FilesPort + "/" + fileName
	client := http.Client{
		Timeout: c.Timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	f, err := os.Create(filepath.Base(fileName))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}
