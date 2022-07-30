package fileserver

import (
	"io"
	"net/http"
	"os"
	"path"

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
	logrus.Debug("GetFileHandler: " + fileName)
	fileNameWithPath := path.Join(s.filesDir, fileName)
	file, err := os.Open(fileNameWithPath)
	if err != nil {
		Error(err, w)
		return
	}
	defer file.Close()
	_, err = io.Copy(w, file)
	if err != nil {
		Error(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
