package server

import (
	"fmt"
	"os"
)

// Info is a function that retunrs server info.
func (s *Server) Info() (string, error) {
	var info string
	hostname, err := os.Hostname()
	info = "id: " + s.id
	if err == nil {
		info += "\nhostname: " + hostname
	}
	info += "\nkeys: " + fmt.Sprint(len(s.kv))
	info += "\nmemory: " + fmt.Sprint(bToMb(s.memUsage)) + " MB"
	info += "\nmemory limit: " + fmt.Sprint(bToMb(s.memLimit)) + " MB"
	info += "\nnodes: " + fmt.Sprint(s.nodes)
	info += "\nreplicas: " + fmt.Sprint(s.replicas)
	return info, nil
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
