package fileserver

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// Serve is a function that starts the file server.
func (s *FileServer) Serve() {
	logrus.Info("Starting file server...")
	r := chi.NewRouter()
	r.Get("/*", s.GetFileHandler)
	r.Post("/*", s.PostFileHandler)
	r.Delete("/*", s.DeleteFileHandler)
	http.ListenAndServe(s.listenHost+":"+s.listenPort, r)
}
