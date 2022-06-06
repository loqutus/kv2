package server

import "fmt"

func (s *Server) Info() (string, error) {
	var info string
	info = "keys: " + fmt.Sprint(len(s.kv))
	return info, nil
}
