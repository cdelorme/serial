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
	l := int64(*out)
	return self.SerializeInt64(&l)
}

func (self *WriteStream) SerializeUint(out *uint) error {
	l := uint64(*out)
	return self.SerializeUint64(&l)
}

func (self *WriteStream) SerializeInt8(out *int8) error {
	return binary.Write(self, ByteOrder, *out)
}

func (self *WriteStream) SerializeInt16(out *int16) error {
	return binary.Write(self, ByteOrder, *out)
}

func (self *WriteStream) SerializeInt32(out *int32) error {
	return binary.Write(self, ByteOrder, *out)
}

func (self *WriteStream) SerializeInt64(out *int64) error {
	return binary.Write(self, ByteOrder, *out)
}

func (self *WriteStream) SerializeUint8(out *uint8) error {
	return binary.Write(self, ByteOrder, *out)
}

func (self *WriteStream) SerializeUint16(out *uint16) error {
	return binary.Write(self, ByteOrder, *out)
}

func (self *WriteStream) SerializeUint32(out *uint32) error {
	return binary.Write(self, ByteOrder, *out)
}

func (self *WriteStream) SerializeUint64(out *uint64) error {
	return binary.Write(self, ByteOrder, *out)
}

func (self *WriteStream) SerializeFloat32(out *float32) error {
	return binary.Write(self, ByteOrder, *out)
}

func (self *WriteStream) SerializeFloat64(out *float64) error {
	return binary.Write(self, ByteOrder, *out)
}
