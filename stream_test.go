package transport

import (
	"testing"
)

func TestStreamString(t *testing.T) {
	t.Parallel()
	out := "bananas"

	// test encoding
	w := WriteStream{}
	if e := w.SerializeString(&out); e != nil || w.Len() == 0 {
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
	r.Write(w.Bytes()[:8])
	if e := r.SerializeString(&in); e == nil {
		t.Logf("%s\n", e)
		t.FailNow()
	}

	// test full decode
	r.Buffer = w.Buffer
	if e := r.SerializeString(&in); e != nil || out != in {
		t.FailNow()
	}

	// test write empty string
	w.Reset()
	var g string
	if e := w.SerializeString(&g); e != nil || w.Len() == 0 {
		t.FailNow()
	}
}

func TestStreamInt(t *testing.T) {
	t.Parallel()
	o := 12

	// test writing an int
	w := WriteStream{}
	if e := w.SerializeInt(&o); e != nil || w.Len() == 0 {
		t.FailNow()
	}

	// test reading an int fail
	var i int
	r := ReadStream{}
	if e := r.SerializeInt(&i); e == nil {
		t.FailNow()
	}

	// test reading an int
	r.Buffer = w.Buffer
	if e := r.SerializeInt(&i); e != nil || i != o {
		t.FailNow()
	}
}
