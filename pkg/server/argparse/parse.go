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
	// Port to listen for fileserver connections.
	listenPortFileserver := os.Getenv("KV2_LISTEN_PORT_FILESERVER")
	if listenPortFileserver == "" {
		listenPortFileserver = "6971"
	}
	// Memory limit.
	memLimit, err := strconv.Atoi(os.Getenv("KV2_MEM_LIMIT"))
	if memLimit == 0 || err != nil {
		memLimit = 1024 * 1024 * 1024
	}
	// number of replicas.
	replicas, err := strconv.Atoi(os.Getenv("KV2_REPLICAS"))
	if replicas == 0 || err != nil {
		replicas = 1
	}
	// nodes list
	var nodes [][]string
	nodesString := os.Getenv("KV2_NODES")
	nodes = parseNodes(nodesString)
	return Args{
		ListenPortClient:     listenPortClient,
		ListenPortServer:     listenPortServer,
		ListenPortFileServer: listenPortFileserver,
		ListenHost:           ListenHost,
		MemLimit:             uint64(memLimit),
		Replicas:             replicas,
		NodeNames:            nodes,
	}
}
