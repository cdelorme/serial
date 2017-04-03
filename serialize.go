package transport

import "encoding/binary"

var byteOrder = binary.LittleEndian

type Writer interface {
	Write([]byte) (int, error)
}

type Reader interface {
	Read([]byte) (int, error)
}

func NewReadSerial(r Reader) *ReadSerial {
	return &ReadSerial{r}
}

func NewWriteSerial(w Writer) *WriteSerial {
	return &WriteSerial{w}
}
