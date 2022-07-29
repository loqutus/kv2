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
	// port to listen for fileserver client connections.
	listenPortFileServer := os.Getenv("KV2_FILESERVER_LISTEN_PORT")
	if listenPortFileServer == "" {
		listenPortFileServer = "6971"
	}
	// Host to listen on.
	ListenHost := os.Getenv("KV2_LISTEN_HOST")
	if ListenHost == "" {
		ListenHost = "127.0.0.1"
	}
	// Memory limit.
	memLimit, err := strconv.Atoi(os.Getenv("KV2_MEMORY_LIMIT"))
	if memLimit == 0 || err != nil {
		memLimit = 1024 * 1024 * 1024
	}
	// number of replicas.
	replicas, err := strconv.Atoi(os.Getenv("KV2_REPLICAS"))
	if replicas == 0 || err != nil {
		replicas = 1
	}
	// debug mode.
	debugStr := os.Getenv("KV2_DEBUG")
	debug := false
	if debugStr == "true" {
		debug = true
	}
	// fileserver directory.
	fileServerDir := os.Getenv("KV2_FILESERVER_DIR")
	if fileServerDir == "" {
		fileServerDir = "/tmp/kv2/fileserver"
	}
	// nodes list
	var nodes [][]string
	nodesString := os.Getenv("KV2_NODES")
	nodes = parseNodes(nodesString)
	return Args{
		ListenPortClient:     listenPortClient,
		ListenPortServer:     listenPortServer,
		ListenPortFileServer: listenPortFileServer,
		ListenHost:           ListenHost,
		MemLimit:             uint64(memLimit),
		Replicas:             replicas,
		NodeNames:            nodes,
		Debug:                debug,
		FileServerDir:        fileServerDir,
	}
}
