package server

import "github.com/rusik69/kv2/pkg/server/argparse"

// Setup the server args.
func SetupArgs() {
	args := argparse.Parse()
	ServerInstance.listenPortClient = args.ListenPortClient
	ServerInstance.listenPortServer = args.ListenPortServer
	ServerInstance.listenHost = args.ListenHost
	return
}
