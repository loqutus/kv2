package argparse

type Args struct {
	ListenPortClient     string
	ListenPortServer     string
	ListenPortFileServer string
	ListenHost           string
	MemLimit             uint64
	Replicas             int
	NodeNames            [][]string
	Debug                bool
	FileServerDir        string
}
