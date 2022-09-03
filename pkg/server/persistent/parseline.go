package persistent

import (
	"bytes"
	"encoding/gob"
)

// parseLine
func (p *Persistent) parseLine(line string) (BinaryData, error) {
	var data BinaryData
	dec := gob.NewDecoder(bytes.NewBufferString(line))
	err := dec.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
