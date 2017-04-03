package transport

import (
	"encoding/binary"
)

type WriteSerial struct {
	w Writer
}

func (w *WriteSerial) Serialize(out ...interface{}) error {
	for i := range out {
		if e := binary.Write(w.w, byteOrder, out[i]); e != nil {
			return e
		}
	}
	return nil
}
