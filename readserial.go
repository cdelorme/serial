package transport

import (
	"encoding/binary"
)

type ReadSerial struct {
	r Reader
}

func (r *ReadSerial) Serialize(in ...interface{}) error {
	for i := range in {
		if e := binary.Read(r.r, byteOrder, in[i]); e != nil {
			return e
		}
	}
	return nil
}
