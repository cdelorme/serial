package transport

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const MaxUint8 = ^uint8(0)
const MaxUint16 = ^uint16(0)
const MaxUint32 = ^uint32(0)
const MaxUint64 = ^uint64(0)

const MaxInt8 = int8(MaxUint8 >> 1)
const MaxInt16 = int16(MaxUint16 >> 1)
const MaxInt32 = int32(MaxUint32 >> 1)
const MaxInt64 = int64(MaxUint64 >> 1)

func NewReadSerial(buf []byte) *ReadSerial {
	return &ReadSerial{bytes.NewBuffer(buf)}
}

func NewWriteSerial(buf []byte) *WriteSerial {
	return &WriteSerial{bytes.NewBuffer(buf)}
}

var ByteOrder = binary.LittleEndian
var MaxSizeExceeded = errors.New("Exceeded maximum parameter size...")
