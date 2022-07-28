package fileserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *FileServer) InfoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	logrus.Info("InfoHandler")
	w.Write([]byte("InfoHandler"))
	return
}
