package transport

import (
	"bytes"
	"encoding/binary"
)

type WriteStream struct {
	bytes.Buffer
}

func (self *WriteStream) SerializeString(out *string, maxSize uint64) error {
	b := []byte(*out)
	l := uint(len(b))

	if e := self.SerializeUint(&l, maxSize); e != nil {
		return e
	}
	return binary.Write(self, ByteOrder, b)
}

func (self *WriteStream) SerializeInt(out *int, maxSize int64) error {
	l := int64(*out)

	if maxSize != 0 && l > maxSize {
		return MaxSizeExceeded
	}

	switch {
	case maxSize == 0 || maxSize > int64(MaxInt32):
		return self.SerializeInt64(&l)
	case maxSize > int64(MaxInt16):
		el := int32(l)
		return self.SerializeInt32(&el)
	case maxSize > int64(MaxInt8):
		el := int16(l)
		return self.SerializeInt16(&el)
	default:
		el := int8(l)
		return self.SerializeInt8(&el)
	}
}

func (self *WriteStream) SerializeUint(out *uint, maxSize uint64) error {
	l := uint64(*out)

	if maxSize > 0 && l > maxSize {
		return MaxSizeExceeded
	}

	switch {
	case maxSize == 0 || maxSize > uint64(MaxUint32):
		return self.SerializeUint64(&l)
	case maxSize > uint64(MaxUint16):
		el := uint32(l)
		return self.SerializeUint32(&el)
	case maxSize > uint64(MaxUint8):
		el := uint16(l)
		return self.SerializeUint16(&el)
	default:
		el := uint8(l)
		return self.SerializeUint8(&el)
	}
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
