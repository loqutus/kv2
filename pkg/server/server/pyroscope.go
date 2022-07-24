package server

import (
	"os"

	"github.com/pyroscope-io/client/pyroscope"
	"github.com/sirupsen/logrus"
)

func startPyroscope() {
	logrus.Println("Starting pyroscope")
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "kv2",
		ServerAddress:   os.Getenv("KV2_PYROSCOPE_SERVER"),
		Logger:          pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
}
