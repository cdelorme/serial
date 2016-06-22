package transport

import (
	"testing"
)

func TestPlacebo(_ *testing.T) {}

func TestEntity(t *testing.T) {
	t.Parallel()
	name := "Casey"

	// serialize data to write stream
	e := Entity{Name: name}
	w := &WriteStream{}
	if err := e.Serialize(w); err != nil || len(w.Data) == 0 {
		t.FailNow()
	}

	// de-serialize from read stream using previous write stream's data
	i := Entity{}
	r := &ReadStream{Data: w.Data}
	if err := i.Serialize(r); err != nil || i.Name != name {
		t.FailNow()
	}

	// force error with invalid data
	r.Reset()
	r.Data = w.Data[:9]
	if err := i.Serialize(r); err == nil {
		t.FailNow()
	}

	// verify that we never successfully parse partial integers
	for n := 0; n < len(w.Data); n++ {
		r.Reset()
		r.Data = w.Data[:n]
		if e := i.Serialize(r); e == nil {
			t.FailNow()
		}
	}
}
