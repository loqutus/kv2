package server

import "github.com/rusik69/kv2/pkg/server/argparse"

// Setup the server args.
func (s *Server) SetupArgs(args argparse.Args) {
	s.listenPortClient = args.ListenPortClient
	s.listenPortServer = args.ListenPortServer
	s.listenHost = args.ListenHost
}
