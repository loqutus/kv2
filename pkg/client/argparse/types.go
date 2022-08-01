package argparse

// Command line args
type Args struct {
	ServerHost      string
	ServerPort      string
	ServerFilesPort string
}

// ArgsInstance is a singleton of Args.
var ArgsInstance Args
