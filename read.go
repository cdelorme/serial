package serial

import (
	"encoding/binary"
)

// A serialization writer which exposes the Serialize function.
type Read struct {
	r Reader
}

// This method directly funnels all input into binary.Write, which only accepts
// fixed size data types, and will not work on variable length data such as int,
// uint, string, slices of those types, and structures that contain them.
//
// It will return immediately on the first error encountered.
//
// While it can accept a slice and populate the existing instances, but it does
// not restore dynamic sized records.  To work around this, expect the size to
// be stored and load that first.
func (r *Read) Serialize(in ...interface{}) error {
	for i := range in {
		if e := binary.Read(r.r, byteOrder, in[i]); e != nil {
			return e
		}
	}
	return nil
}
