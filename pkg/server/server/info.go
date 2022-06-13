package server

import (
	"fmt"
	"runtime"
)

// Info is a function that retunrs server info.
func (s *Server) Info() (string, error) {
	var info string
	info = "keys: " + fmt.Sprint(len(s.kv))
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	info += "\nmemory: " + fmt.Sprint(bToMb(m.Alloc)) + " MB"
	info += "\nnodes: " + fmt.Sprint(len(s.nodes))
	info += "\nreplicas: " + fmt.Sprint(s.replicas)
	return info, nil
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
