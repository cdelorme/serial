package transport

import (
	"testing"
)

func TestStreamString(t *testing.T) {
	t.Parallel()
	out := "bananas"

	// test encoding
	w := WriteStream{}
	if e := w.SerializeString(&out); e != nil || len(w.Data) == 0 {
		t.FailNow()
	}

	// test decode empty error
	var in string
	r := ReadStream{}
	if e := r.SerializeString(&in); e == nil {
		t.FailNow()
	}

	// test decode partial error
	r.Reset()
	r.Data = w.Data[:8]
	if e := r.SerializeString(&in); e == nil {
		t.FailNow()
	}

	// test full decode
	r.Reset()
	r.Data = w.Data
	if e := r.SerializeString(&in); e != nil || out != in {
		t.FailNow()
	}

	// test write empty string
	w.Data = []byte{}
	var g string
	if e := w.SerializeString(&g); e != nil || len(w.Data) == 0 {
		t.FailNow()
	}
}

func TestStreamInt(t *testing.T) {
	t.Parallel()
	o := 12

	// test writing an int
	w := WriteStream{}
	if e := w.SerializeInt(&o); e != nil || len(w.Data) == 0 {
		t.FailNow()
	}

	// test reading an int fail
	var i int
	r := ReadStream{}
	if e := r.SerializeInt(&i); e == nil {
		t.FailNow()
	}

	// test reading an int
	r.Reset()
	r.Data = w.Data
	if e := r.SerializeInt(&i); e != nil || i != o {
		t.Logf("%s\n", e)
		t.FailNow()
	}
}
