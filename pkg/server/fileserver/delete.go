package fileserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// DeleteFileHandler is a function that handles DELETE requests.
func (s *FileServer) DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if len(r.URL.Path) < 2 {
		logrus.Info("DeleteFileHandler: file name is not specified")
		w.Write([]byte("file name is not specified"))
		r.Response.StatusCode = http.StatusBadRequest
		return
	}
	fileName := r.URL.Path[1:]
	logrus.Info("DeleteFileHandler: " + fileName)
	defer r.Body.Close()
	return
}
