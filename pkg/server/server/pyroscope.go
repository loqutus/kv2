package server

import (
	"os"

	pyroscope "github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"github.com/sirupsen/logrus"
)

func startPyroscope() {
	logrus.Println("Starting Pyroscope")
	pyroscope.Start(pyroscope.Config{
		ApplicationName: os.Getenv("kv2"),
		ServerAddress:   os.Getenv("KV2_PYROSCOPE_SERVER"),
	})
}
