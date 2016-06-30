package transport

import (
	"bytes"
	"testing"
)

func TestEntity(t *testing.T) {
	t.Parallel()
	name := "Casey"
	var b bytes.Buffer
	o, i := Entity{Name: name, Health: [2]uint16{100, 1000}}, Entity{}
	r, w := &ReadStream{Buffer: &bytes.Buffer{}}, &WriteStream{Buffer: &b}

	// serialize data to write stream
	if err := o.Serialize(w); err != nil || w.Len() == 0 {
		t.FailNow()
	}

	// force error with invalid data
	r.Write(w.Bytes()[:8])
	if err := i.Serialize(r); err == nil {
		t.FailNow()
	}

	// verify that we never successfully parse partial integers
	for n := 0; n < w.Len(); n++ {
		r.Reset()
		r.Write(w.Bytes()[:n])
		if e := i.Serialize(r); e == nil {
			t.FailNow()
		}
	}

	// de-serialize from read stream using previous write stream's data
	r.Buffer = w.Buffer
	if err := i.Serialize(r); err != nil || i.Name != name {
		t.FailNow()
	}
}
