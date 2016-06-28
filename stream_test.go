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

func TestStreamUint(t *testing.T) {
	t.Parallel()
	var o uint = 12

	w := WriteStream{}
	if e := w.SerializeUint(&o); e != nil || w.Len() == 0 {
		t.FailNow()
	}

	var i uint
	r := ReadStream{}
	if e := r.SerializeUint(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeUint(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestStreamInt8(t *testing.T) {
	t.Parallel()
	var o int8 = 12

	w := WriteStream{}
	if e := w.SerializeInt8(&o); e != nil || w.Len() != 1 {
		t.FailNow()
	}

	var i int8
	r := ReadStream{}
	if e := r.SerializeInt8(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeInt8(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestStreamInt16(t *testing.T) {
	t.Parallel()
	var o int16 = 12

	w := WriteStream{}
	if e := w.SerializeInt16(&o); e != nil || w.Len() != 2 {
		t.FailNow()
	}

	var i int16
	r := ReadStream{}
	if e := r.SerializeInt16(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeInt16(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestStreamInt32(t *testing.T) {
	t.Parallel()
	var o int32 = 12

	w := WriteStream{}
	if e := w.SerializeInt32(&o); e != nil || w.Len() != 4 {
		t.FailNow()
	}

	var i int32
	r := ReadStream{}
	if e := r.SerializeInt32(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeInt32(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestStreamInt64(t *testing.T) {
	t.Parallel()
	var o int64 = 12

	w := WriteStream{}
	if e := w.SerializeInt64(&o); e != nil || w.Len() != 8 {
		t.FailNow()
	}

	var i int64
	r := ReadStream{}
	if e := r.SerializeInt64(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeInt64(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestStreamUint8(t *testing.T) {
	t.Parallel()
	var o uint8 = 12

	w := WriteStream{}
	if e := w.SerializeUint8(&o); e != nil || w.Len() != 1 {
		t.FailNow()
	}

	var i uint8
	r := ReadStream{}
	if e := r.SerializeUint8(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeUint8(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestStreamUint16(t *testing.T) {
	t.Parallel()
	var o uint16 = 12

	w := WriteStream{}
	if e := w.SerializeUint16(&o); e != nil || w.Len() != 2 {
		t.FailNow()
	}

	var i uint16
	r := ReadStream{}
	if e := r.SerializeUint16(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeUint16(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestStreamUint32(t *testing.T) {
	t.Parallel()
	var o uint32 = 12

	w := WriteStream{}
	if e := w.SerializeUint32(&o); e != nil || w.Len() != 4 {
		t.FailNow()
	}

	var i uint32
	r := ReadStream{}
	if e := r.SerializeUint32(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeUint32(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestStreamUint64(t *testing.T) {
	t.Parallel()
	var o uint64 = 12

	w := WriteStream{}
	if e := w.SerializeUint64(&o); e != nil || w.Len() != 8 {
		t.FailNow()
	}

	var i uint64
	r := ReadStream{}
	if e := r.SerializeUint64(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeUint64(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestStreamFloat32(t *testing.T) {
	t.Parallel()
	var o float32 = 12.3

	w := WriteStream{}
	if e := w.SerializeFloat32(&o); e != nil || w.Len() != 4 {
		t.FailNow()
	}

	var i float32
	r := ReadStream{}
	if e := r.SerializeFloat32(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeFloat32(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestStreamFloat64(t *testing.T) {
	t.Parallel()
	var o float64 = 12.3

	w := WriteStream{}
	if e := w.SerializeFloat64(&o); e != nil || w.Len() != 8 {
		t.FailNow()
	}

	var i float64
	r := ReadStream{}
	if e := r.SerializeFloat64(&i); e == nil {
		t.FailNow()
	}

	r.Buffer = w.Buffer
	if e := r.SerializeFloat64(&i); e != nil || i != o {
		t.FailNow()
	}
}
