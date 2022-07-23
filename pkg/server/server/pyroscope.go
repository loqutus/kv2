package server

import (
	"os"

	pyroscope "github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
)

func startPyroscope() {
	pyroscope.Start(pyroscope.Config{
		ApplicationName: os.Getenv("kv2"),
		ServerAddress:   os.Getenv("KV2_PYROSCOPE_SERVER"),
	})
}
