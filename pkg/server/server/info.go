package server

import (
	"fmt"
	"runtime"
)

// Info is a function that retunrs server info.
func (s *Server) Info() (string, error) {
	var info string
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	info = "id: " + s.id
	info += "\nkeys: " + fmt.Sprint(len(s.kv))
	info += "\nmemory: " + fmt.Sprint(bToMb(m.Alloc)) + " MB"
	info += "\nnodes: " + fmt.Sprint(s.nodes)
	info += "\nreplicas: " + fmt.Sprint(s.replicas)
	return info, nil
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
