package transport

import (
	"bytes"
	"encoding/binary"
)

type ReadStream struct {
	bytes.Buffer
}

func (self *ReadStream) SerializeString(in *string) error {
	var l int
	if e := self.SerializeInt(&l); e != nil {
		return e
	}

	d := make([]byte, l)
	if _, e := self.Read(d); e != nil {
		return e
	}
	*in = string(d)

	return nil
}

func (self *ReadStream) SerializeInt(in *int) error {
	var l int64
	if e := binary.Read(self, ByteOrder, &l); e != nil {
		return e
	}

	*in = int(l)

	return nil
}
