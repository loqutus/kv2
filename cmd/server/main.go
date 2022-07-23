package main

import (
	"github.com/rusik69/kv2/pkg/server/argparse"
	"github.com/rusik69/kv2/pkg/server/logger"
	"github.com/rusik69/kv2/pkg/server/server"
)

// main server cmd function.
func main() {
	logger.SetLoggerFormat()
	args := argparse.Parse()
	server.ServerInstance.SetupArgs(args)
	go server.ServerInstance.MemController()
	go server.ServerInstance.ConnController()
	server.ServerInstance.Serve()
}
