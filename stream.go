package transport

import "encoding/binary"

type Stream interface {
	SerializeString(*string) error
	SerializeInt(*int) error
}

var ByteOrder = binary.LittleEndian
