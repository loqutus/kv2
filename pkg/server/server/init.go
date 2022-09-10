package server

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/rusik69/kv2/pkg/server/argparse"
	"github.com/rusik69/kv2/pkg/server/persistent"
)

// init is a function that initializes the server.
func (s *Server) Init(args argparse.Args) error {
	s.kv = make(map[string]int64)
	s.id = s.getID()
	s.replicas = args.Replicas
	persistent.PersistentInstance.Init(args, s.kv)
	return nil
}

// getID is a function that gets the server id from the file.
func (s *Server) getID() string {
	fileName := path.Join(s.stateDir, "id")
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return uuid.New().String()
	} else {
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			return uuid.New().String()
		}
		return string(fileBytes)
	}
}
