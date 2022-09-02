package persistent

import (
	"bufio"
	"io"
)

// readKV is a function that reads a key-value pair from the persistent storage.
func (p *Persistent) readKV(kv map[string]int64) error {
	var offset int64
	scanner := bufio.NewScanner(p.F)
	for scanner.Scan() {
		key, err := p.parseLine(scanner.Text())
		if err != nil {
			continue
		}
		kv[string(key)] = offset
		currentOffset, err := p.F.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
		offset = currentOffset
	}
	return nil
}
