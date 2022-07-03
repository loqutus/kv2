package server

import (
	"runtime"
	"time"
)

// MemController is a function that controls the memory usage of the server.
func (s *Server) MemController() {
	var m runtime.MemStats
	for {
		runtime.ReadMemStats(&m)
		s.memUsage = m.Alloc
		time.Sleep(time.Second)
	}
}
