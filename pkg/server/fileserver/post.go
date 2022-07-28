package fileserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// PostFileHandler is a function that handles POST requests.
func (s *FileServer) PostFileHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if len(r.URL.Path) < 2 {
		logrus.Info("PostFileHandler: file name is not specified")
		w.Write([]byte("file name is not specified"))
		r.Response.StatusCode = http.StatusBadRequest
		return
	}
	fileName := r.URL.Path[1:]
	logrus.Info("PostFileHandler: " + fileName)
	defer r.Body.Close()
	return
}
