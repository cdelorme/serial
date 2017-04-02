package transport

import (
	"bytes"
	"encoding/binary"
)

type WriteSerial struct {
	*bytes.Buffer
}

func (w *WriteSerial) SerializeString(out *string, maxSize uint64) error {
	b := []byte(*out)
	l := uint(len(b))

	if e := w.SerializeUint(&l, maxSize); e != nil {
		return e
	}
	return binary.Write(w, ByteOrder, b)
}

func (w *WriteSerial) SerializeBool(out *bool) error {
	var t uint8
	if *out {
		t = 1
	}
	return w.SerializeUint8(&t)
}

func (w *WriteSerial) SerializeInt(out *int, maxSize int64) error {
	l := int64(*out)

	if maxSize != 0 && l > maxSize {
		return MaxSizeExceeded
	}

	switch {
	case maxSize == 0 || maxSize > int64(MaxInt32):
		return w.SerializeInt64(&l)
	case maxSize > int64(MaxInt16):
		el := int32(l)
		return w.SerializeInt32(&el)
	case maxSize > int64(MaxInt8):
		el := int16(l)
		return w.SerializeInt16(&el)
	default:
		el := int8(l)
		return w.SerializeInt8(&el)
	}
}

func (w *WriteSerial) SerializeUint(out *uint, maxSize uint64) error {
	l := uint64(*out)

	if maxSize > 0 && l > maxSize {
		return MaxSizeExceeded
	}

	switch {
	case maxSize == 0 || maxSize > uint64(MaxUint32):
		return w.SerializeUint64(&l)
	case maxSize > uint64(MaxUint16):
		el := uint32(l)
		return w.SerializeUint32(&el)
	case maxSize > uint64(MaxUint8):
		el := uint16(l)
		return w.SerializeUint16(&el)
	default:
		el := uint8(l)
		return w.SerializeUint8(&el)
	}
}

func (w *WriteSerial) SerializeInt8(out *int8) error {
	return binary.Write(w, ByteOrder, *out)
}

func (w *WriteSerial) SerializeInt16(out *int16) error {
	return binary.Write(w, ByteOrder, *out)
}

func (w *WriteSerial) SerializeInt32(out *int32) error {
	return binary.Write(w, ByteOrder, *out)
}

func (w *WriteSerial) SerializeInt64(out *int64) error {
	return binary.Write(w, ByteOrder, *out)
}

func (w *WriteSerial) SerializeUint8(out *uint8) error {
	return binary.Write(w, ByteOrder, *out)
}

func (w *WriteSerial) SerializeUint16(out *uint16) error {
	return binary.Write(w, ByteOrder, *out)
}

func (w *WriteSerial) SerializeUint32(out *uint32) error {
	return binary.Write(w, ByteOrder, *out)
}

func (w *WriteSerial) SerializeUint64(out *uint64) error {
	return binary.Write(w, ByteOrder, *out)
}

func (w *WriteSerial) SerializeFloat32(out *float32) error {
	return binary.Write(w, ByteOrder, *out)
}

func (w *WriteSerial) SerializeFloat64(out *float64) error {
	return binary.Write(w, ByteOrder, *out)
}
