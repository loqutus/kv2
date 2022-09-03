package persistent

import "bufio"

// Get is a function that gets a value by offset from the persistent storage.
func (p *Persistent) Get(offset int64) ([]byte, error) {
	_, err := p.F.Seek(offset, 0)
	if err != nil {
		return []byte{}, err
	}
	reader := bufio.NewReader(p.F)
	line, _, err := reader.ReadLine()
	if err != nil {
		return []byte{}, err
	}
	data, err := p.parseLine(string(line))
	if err != nil {
		return []byte{}, err
	}
	_, err = p.F.Seek(p.offset, 0)
	if err != nil {
		return []byte{}, err
	}
	return data.Val, nil
}
