package transport

import (
	"bytes"
	"encoding/binary"
)

type WriteStream struct {
	bytes.Buffer
}

func (self *WriteStream) SerializeString(out *string) error {
	b := []byte(*out)
	l := len(b)
	self.SerializeInt(&l)
	return binary.Write(self, ByteOrder, b)
}

func (self *WriteStream) SerializeInt(out *int) error {
	return binary.Write(self, ByteOrder, int64(*out))
}
