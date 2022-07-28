package fileserver

import "github.com/rusik69/kv2/pkg/server/argparse"

// SetupArgs is a function that sets up the arguments for the fileserver.
func (s *FileServer) SetupArgs(args argparse.Args) {
	s.listenHost = args.ListenHost
	s.listenPort = args.ListenPortFileServer
	s.filesDir = args.FileServerDir
	s.replicas = args.Replicas
	return
}
