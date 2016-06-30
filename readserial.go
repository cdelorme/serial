package transport

import (
	"bytes"
	"encoding/binary"
)

type ReadSerial struct {
	*bytes.Buffer
}

func (self *ReadSerial) SerializeString(in *string, maxSize uint64) error {
	var l uint
	if e := self.SerializeUint(&l, maxSize); e != nil {
		return e
	}

	d := make([]byte, l)
	if _, e := self.Read(d); e != nil {
		return e
	}
	*in = string(d)

	return nil
}

func (self *ReadSerial) SerializeInt(in *int, maxSize int64) error {
	switch {
	case maxSize == 0 || maxSize > int64(MaxInt32):
		var l int64
		if e := self.SerializeInt64(&l); e != nil {
			return e
		}
		if maxSize != 0 && l > maxSize {
			return MaxSizeExceeded
		}
		*in = int(l)
	case maxSize > int64(MaxInt16):
		var l int32
		if e := self.SerializeInt32(&l); e != nil {
			return e
		}
		if maxSize != 0 && int64(l) > maxSize {
			return MaxSizeExceeded
		}
		*in = int(l)
	case maxSize > int64(MaxInt8):
		var l int16
		if e := self.SerializeInt16(&l); e != nil {
			return e
		}
		if maxSize != 0 && int64(l) > maxSize {
			return MaxSizeExceeded
		}
		*in = int(l)
	default:
		var l int8
		if e := self.SerializeInt8(&l); e != nil {
			return e
		}
		if maxSize != 0 && int64(l) > maxSize {
			return MaxSizeExceeded
		}
		*in = int(l)
	}

	return nil
}

func (self *ReadSerial) SerializeUint(in *uint, maxSize uint64) error {
	switch {
	case maxSize == 0 || maxSize > uint64(MaxUint32):
		var l uint64
		if e := self.SerializeUint64(&l); e != nil {
			return e
		}
		if maxSize > 0 && l > maxSize {
			return MaxSizeExceeded
		}
		*in = uint(l)
	case maxSize > uint64(MaxUint16):
		var l uint32
		if e := self.SerializeUint32(&l); e != nil {
			return e
		}
		if maxSize > 0 && uint64(l) > maxSize {
			return MaxSizeExceeded
		}
		*in = uint(l)
	case maxSize > uint64(MaxUint8):
		var l uint16
		if e := self.SerializeUint16(&l); e != nil {
			return e
		}
		if maxSize > 0 && uint64(l) > maxSize {
			return MaxSizeExceeded
		}
		*in = uint(l)
	default:
		var l uint8
		if e := self.SerializeUint8(&l); e != nil {
			return e
		}
		if maxSize > 0 && uint64(l) > maxSize {
			return MaxSizeExceeded
		}
		*in = uint(l)
	}

	return nil
}

func (self *ReadSerial) SerializeInt8(in *int8) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadSerial) SerializeInt16(in *int16) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadSerial) SerializeInt32(in *int32) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadSerial) SerializeInt64(in *int64) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadSerial) SerializeUint8(in *uint8) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadSerial) SerializeUint16(in *uint16) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadSerial) SerializeUint32(in *uint32) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadSerial) SerializeUint64(in *uint64) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadSerial) SerializeFloat32(in *float32) error {
	return binary.Read(self, ByteOrder, in)
}

func (self *ReadSerial) SerializeFloat64(in *float64) error {
	return binary.Read(self, ByteOrder, in)
}
