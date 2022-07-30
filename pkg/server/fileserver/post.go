package fileserver

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// PostFileHandler is a function that handles POST requests.
func (s *FileServer) PostFileHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if len(r.URL.Path) < 2 {
		Error(errors.New("file name is not specified"), w)
		return
	}
	fileName := r.URL.Path[1:]
	fileNameWithPath := path.Join(s.filesDir, fileName)
	logrus.Debug("PostFileHandler: " + fileName)
	err := s.EnsureDir(fileNameWithPath)
	if err != nil {
		Error(err, w)
		return
	}
	file, err := os.Create(fileNameWithPath)
	if err != nil {
		Error(err, w)
		return
	}
	_, err = io.Copy(file, r.Body)
	if err != nil {
		Error(err, w)
		return
	}
	return
}
