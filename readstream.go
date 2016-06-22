package transport

import (
	"encoding/binary"
	"io"
)

type ReadStream struct {
	position int
	Data     []byte
}

func (self *ReadStream) SerializeString(in *string) error {
	var l int64
	if e := binary.Read(self, binary.LittleEndian, &l); e != nil {
		return e
	}

	if self.position+int(l)-1 >= len(self.Data) {
		return io.EOF
	}

	*in = string(self.Data[self.position : self.position+int(l)])
	self.position += int(l)

	return nil
}

func (self *ReadStream) SerializeInt(in *int) error {
	var l int64
	if e := binary.Read(self, binary.LittleEndian, &l); e != nil {
		return e
	}

	*in = int(l)

	return nil
}

func (self *ReadStream) Read(p []byte) (int, error) {
	if self.position >= len(self.Data) {
		return 0, io.EOF
	}
	p[0] = self.Data[self.position]
	self.position++
	return 1, nil
}

func (self *ReadStream) Reset() {
	self.position = 0
}
