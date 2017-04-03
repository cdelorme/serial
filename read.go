package serial

import (
	"encoding/binary"
)

type Read struct {
	r Reader
}

func (r *Read) Serialize(in ...interface{}) error {
	for i := range in {
		if e := binary.Read(r.r, byteOrder, in[i]); e != nil {
			return e
		}
	}
	return nil
}
