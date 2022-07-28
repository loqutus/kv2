package fileserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetFileHandler is a function that handles GET requests.
func (s *FileServer) GetFileHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if len(r.URL.Path) < 2 {
		logrus.Info("GetFileHandler: file name is not specified")
		w.Write([]byte("file name is not specified"))
		r.Response.StatusCode = http.StatusBadRequest
		return
	}
	fileName := r.URL.Path[1:]
	logrus.Info("GetFileHandler: " + fileName)
	return
}
