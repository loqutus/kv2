package persistent

import (
	"bytes"
	"encoding/gob"
)

func (p *Persistent) parseLine(line string) ([]byte, error) {
	var data BinaryData
	dec := gob.NewDecoder(bytes.NewBufferString(line))
	err := dec.Decode(&data)
	if err != nil {
		return []byte{}, err
	}
	return data.Key, nil
}
