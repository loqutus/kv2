package server

import (
	"os"

	pyroscope "github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"github.com/sirupsen/logrus"
)

func startPyroscope() {
	logrus.Println("Starting pyroscope")
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "kv2",
		ServerAddress:   os.Getenv("KV2_PYROSCOPE_SERVER"),
		Logger:          pyroscope.StandardLogger,
	})
}
