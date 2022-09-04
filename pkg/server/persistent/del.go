package persistent

import (
	"bufio"
	"bytes"
	"encoding/gob"
)

// Del is a function that marks a key-value pair for deletion in the persistent storage.
func (p *Persistent) Del(offset int64) error {
	_, err := p.F.Seek(offset, 0)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(p.F)
	line, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	data, err := p.parseLine(string(line))
	if err != nil {
		return err
	}
	data.Del = true
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(data)
	if err != nil {
		return err
	}
	_, err = p.F.WriteAt(buf.Bytes(), offset)
	if err != nil {
		return err
	}
	return nil
}
