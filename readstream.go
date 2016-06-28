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

func (self *ReadStream) SerializeInt(in *int) (e error) {
	var l int64
	e = self.SerializeInt64(&l)
	*in = int(l)
	return
}

func (self *ReadStream) SerializeUint(in *uint) (e error) {
	var l uint64
	e = self.SerializeUint64(&l)
	*in = uint(l)
	return
}

func (self *ReadStream) SerializeInt8(in *int8) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadStream) SerializeInt16(in *int16) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadStream) SerializeInt32(in *int32) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadStream) SerializeInt64(in *int64) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadStream) SerializeUint8(in *uint8) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadStream) SerializeUint16(in *uint16) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadStream) SerializeUint32(in *uint32) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadStream) SerializeUint64(in *uint64) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadStream) SerializeFloat32(in *float32) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadStream) SerializeFloat64(in *float64) error {
	return binary.Read(self, ByteOrder, in)
}
