package main

import (
	"github.com/rusik69/kv2/pkg/server/argparse"
	"github.com/rusik69/kv2/pkg/server/fileserver"
	"github.com/rusik69/kv2/pkg/server/logger"
	"github.com/rusik69/kv2/pkg/server/persistent"
	"github.com/rusik69/kv2/pkg/server/server"
)

// main server cmd function.
func main() {
	logger.SetLoggerFormat()
	args := argparse.Parse()
	server.ServerInstance.SetupArgs(args)
	fileserver.FileServerInstance.SetupArgs(args)
	server.ServerInstance.Init(args)
	defer persistent.PersistentInstance.F.Close()
	go server.ServerInstance.MemController()
	go server.ServerInstance.ConnController()
	go fileserver.FileServerInstance.Serve()
	server.ServerInstance.Serve()
}
