package serial

import "encoding/binary"

var byteOrder = binary.LittleEndian

// A replica of the io.Writer interface for compatibility minus the import.
type Writer interface {
	Write([]byte) (int, error)
}

// A replica of the io.Reader interface for compatibility minus the import.
type Reader interface {
	Read([]byte) (int, error)
}

// Returns a new serialization reader with functionality matching the writer.
func NewReader(r Reader) *Read {
	return &Read{r}
}

// Returns a new serialization writer with functionality matching the reader.
func NewWriter(w Writer) *Write {
	return &Write{w}
}
