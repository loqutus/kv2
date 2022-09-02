package persistent

import (
	"os"
)

// BinaryData is a type that represents a binary data.
type BinaryData struct {
	Key []byte
	Val []byte
	Del bool
}

// Persistent is a struct that holds persistent data.
type Persistent struct {
	Dir      string
	Filename string
	F        *os.File
	offset   int64
}

// PersistentInstance is a global persistent instance.
var PersistentInstance *Persistent
