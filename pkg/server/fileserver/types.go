package fileserver

// FileServer is a struct that contains the necessary information to serve files.
type FileServer struct {
	listenHost string
	listenPort string
	filesDir   string
	diskFree   uint64
	replicas   int
	rodes      []Node
}

// Node is a struct that contains the necessary information to connect to a node.
type Node struct {
	host string
	port string
}

// FileServerInstance is the main file server instance.
var FileServerInstance FileServer
