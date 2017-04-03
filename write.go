package serial

import (
	"encoding/binary"
)

type Write struct {
	w Writer
}

func (w *Write) Serialize(out ...interface{}) error {
	for i := range out {
		if e := binary.Write(w.w, byteOrder, out[i]); e != nil {
			return e
		}
	}
	return nil
}
