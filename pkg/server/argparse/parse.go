package argparse

import (
	"os"
	"strconv"
)

// Parse parses the command line arguments.
func Parse() Args {
	// Port to listen for client connections.
	listenPortClient := os.Getenv("KV2_LISTEN_PORT_CLIENT")
	if listenPortClient == "" {
		listenPortClient = "6969"
	}
	// Port to listen for server connections.
	listenPortServer := os.Getenv("KV2_LISTEN_PORT_SERVER")
	if listenPortServer == "" {
		listenPortServer = "6970"
	}
	// Host to listen on.
	ListenHost := os.Getenv("KV2_LISTEN_HOST")
	if ListenHost == "" {
		ListenHost = "127.0.0.1"
	}
	// Memory limit.
	memLimit, err := strconv.Atoi(os.Getenv("KV2_MEM_LIMIT"))
	if memLimit == 0 || err != nil {
		memLimit = 1024 * 1024 * 1024
	}
	return Args{
		ListenPortClient: listenPortClient,
		ListenPortServer: listenPortServer,
		ListenHost:       ListenHost,
		MemLimit:         memLimit,
	}
}
