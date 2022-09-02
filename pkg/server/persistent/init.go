package persistent

import (
	"os"
	"path"

	"github.com/rusik69/kv2/pkg/server/argparse"
)

// persistent file format:
// length BinaryData\n

// Init is a function that reads persistent data from the file.
func (p *Persistent) Init(args argparse.Args, kv map[string]int64) error {
	p.Dir = args.StateDir
	p.Filename = args.StateFile
	var err error
	p.F, err = os.OpenFile(path.Join(p.Dir, p.Filename), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	err = p.readKV(kv)
	if err != nil {
		return err
	} else {
		return nil
	}
}
