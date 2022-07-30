package fileserver

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

func (s *FileServer) EnsureDir(fileName string) error {
	dir := path.Dir(fileName)
	if dir == "." {
		return nil
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logrus.Debug("Creating directory:", dir)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
