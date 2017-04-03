package serial

import (
	"encoding/binary"
)

// A serialization writer which exposes the Serialize function.
type Write struct {
	w Writer
}

// This method directly funnels all input into binary.Write, which only accepts
// fixed size data types, and will not work on variable length data such as int,
// uint, string, slices of those types, and structures that contain them.
//
// It will return immediately on the first error encountered.
//
// While slices of fixed-size types are accepted, the size of that slice is not
// stored with the data.  To work around this, store the size first.
func (w *Write) Serialize(out ...interface{}) error {
	for i := range out {
		if e := binary.Write(w.w, byteOrder, out[i]); e != nil {
			return e
		}
	}
	return nil
}
