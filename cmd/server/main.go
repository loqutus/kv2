package main

import (
	"github.com/rusik69/kv2/pkg/server/logger"
	"github.com/rusik69/kv2/pkg/server/server"
)

// main server cmd function.
func main() {
	logger.SetLoggerFormat()
	server.SetupArgs()
	server.ServerInstance.Listen()
	server.ServerInstance.Serve()
}
