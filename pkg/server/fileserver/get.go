package fileserver

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetFileHandler is a function that handles GET requests.
func (s *FileServer) GetFileHandler(w http.ResponseWriter, r *http.Request) error {
	if len(r.URL.Path) < 2 {
		return errors.New("file name is not specified")
	}
	fileName := r.URL.Path[1:]
	logrus.Info("GetFileHandler: " + fileName)
	defer r.Body.Close()
	
	return nil
}
