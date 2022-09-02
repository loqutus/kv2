package persistent

import (
	"bytes"
	"encoding/gob"
)

// Set is a function that sets a key-value pair in the file
func (p *Persistent) Set(data BinaryData) (int64, error) {
	off := p.offset
	currentOffset := p.offset
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return currentOffset, err
	}
	_, err = p.F.WriteAt(buf.Bytes(), currentOffset)
	if err != nil {
		return 0, err
	}
	currentOffset += int64(len(buf.Bytes()))
	_, err = p.F.WriteAt([]byte("\n"), currentOffset)
	if err != nil {
		return 0, err
	}
	p.offset = currentOffset + 1
	return off, nil
}
