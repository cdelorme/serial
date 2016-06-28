package transport

import (
	"testing"
)

func TestPlacebo(_ *testing.T) {}

func TestEntity(t *testing.T) {
	t.Parallel()
	name := "Casey"

	// serialize data to write stream
	e := Entity{Name: name, Health: [2]int{100, 1000}}
	w := &WriteStream{}
	if err := e.Serialize(w); err != nil || w.Len() == 0 {
		t.FailNow()
	}

	// de-serialize from read stream using previous write stream's data
	i := Entity{}
	r := &ReadStream{Buffer: w.Buffer}
	if err := i.Serialize(r); err != nil || i.Name != name {
		t.FailNow()
	}

	// force error with invalid data
	r.Reset()
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
}
