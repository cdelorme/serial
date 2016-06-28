package transport

import (
	"testing"
)

func TestStreamString(t *testing.T) {
	t.Parallel()
	out := "bananas"

	// test encoding (default)
	w := WriteStream{}
	if e := w.SerializeString(&out, 0); e != nil || w.Len() == 0 {
		t.FailNow()
	}

	// test optimized encoding
	ow := WriteStream{}
	if e := ow.SerializeString(&out, 255); e != nil || w.Len() <= ow.Len() {
		t.FailNow()
	}

	// test max length safety
	if e := ow.SerializeString(&out, 2); e == nil {
		t.FailNow()
	}

	// test decode empty error
	var in string
	r := ReadStream{}
	if e := r.SerializeString(&in, 0); e == nil {
		t.FailNow()
	}

	// test decode partial error
	r.Reset()
	r.Write(w.Bytes()[:8])
	if e := r.SerializeString(&in, 0); e == nil {
		t.FailNow()
	}

	// test full decode
	r.Buffer = w.Buffer
	if e := r.SerializeString(&in, 0); e != nil || out != in {
		t.FailNow()
	}

	// test write empty string
	w.Reset()
	var g string
	if e := w.SerializeString(&g, 0); e != nil || w.Len() == 0 {
		t.FailNow()
	}
}

func TestStreamInt(t *testing.T) {
	t.Parallel()
	r, w := ReadStream{}, WriteStream{}
	var o, i int

	// test read int64 no data
	if e := r.SerializeInt(&i, MaxInt64); e == nil {
		t.FailNow()
	}

	// test read int32 no data
	if e := r.SerializeInt(&i, int64(MaxInt32)); e == nil {
		t.FailNow()
	}

	// test read int16 no data
	if e := r.SerializeInt(&i, int64(MaxInt16)); e == nil {
		t.FailNow()
	}

	// test read int8 no data
	if e := r.SerializeInt(&i, int64(MaxInt8)); e == nil {
		t.FailNow()
	}

	// test write exceeds max
	o = 2
	if e := w.SerializeInt(&o, 1); e == nil {
		t.FailNow()
	}

	// test write int64
	w.Reset()
	o = int(MaxInt64 - 1)
	if e := w.SerializeInt(&o, MaxInt64); e != nil || w.Len() != 8 {
		t.Failed()
	}

	// test read int64 exceeds max
	r.Buffer = w.Buffer
	if e := r.SerializeInt(&i, MaxInt64-2); e == nil {
		t.FailNow()
	}

	// test read int64
	r.Buffer = w.Buffer
	if e := r.SerializeInt(&i, MaxInt64); e != nil || i != o {
		t.FailNow()
	}

	// test write int32
	w.Reset()
	o = int(MaxInt32 - 1)
	t.Logf("Storing: %d\n", o)
	if e := w.SerializeInt(&o, int64(MaxInt32)); e != nil || w.Len() != 4 {
		t.FailNow()
	}

	// test read int32 exceeds max
	r.Buffer = w.Buffer
	if e := r.SerializeInt(&i, int64(MaxInt32-2)); e == nil {
		t.FailNow()
	}

	// test read int32
	r.Buffer = w.Buffer
	if e := r.SerializeInt(&i, int64(MaxInt32)); e != nil {
		t.FailNow()
	}

	// test write int16
	w.Reset()
	o = int(MaxInt16 - 1)
	if e := w.SerializeInt(&o, int64(MaxInt16)); e != nil || w.Len() != 2 {
		t.FailNow()
	}

	// test read int16 exceeds max
	r.Buffer = w.Buffer
	if e := r.SerializeInt(&i, int64(MaxInt16-2)); e == nil {
		t.FailNow()
	}

	// test read int16
	r.Buffer = w.Buffer
	if e := r.SerializeInt(&i, int64(MaxInt16)); e != nil {
		t.FailNow()
	}

	// test write int8
	w.Reset()
	o = int(MaxInt8 - 1)
	if e := w.SerializeInt(&o, int64(MaxInt8)); e != nil || w.Len() != 1 {
		t.FailNow()
	}

	// test read int8 exceeds max
	r.Buffer = w.Buffer
	if e := r.SerializeInt(&i, int64(MaxInt8-2)); e == nil {
		t.FailNow()
	}

	// test read int8
	r.Buffer = w.Buffer
	if e := r.SerializeInt(&i, int64(MaxInt8)); e != nil {
		t.FailNow()
	}
}

func TestStreamUint(t *testing.T) {
	t.Parallel()
	r, w := ReadStream{}, WriteStream{}
	var o, i uint

	// test read uint64 no data
	if e := r.SerializeUint(&i, MaxUint64); e == nil {
		t.FailNow()
	}

	// test read uint32 no data
	if e := r.SerializeUint(&i, uint64(MaxUint32)); e == nil {
		t.FailNow()
	}

	// test read uint16 no data
	if e := r.SerializeUint(&i, uint64(MaxUint16)); e == nil {
		t.FailNow()
	}

	// test read uint8 no data
	if e := r.SerializeUint(&i, uint64(MaxUint8)); e == nil {
		t.FailNow()
	}

	// test write exceeds max
	o = 2
	if e := w.SerializeUint(&o, 1); e == nil {
		t.FailNow()
	}

	// test write uint64
	o = uint(MaxUint64 - 1)
	if e := w.SerializeUint(&o, MaxUint64); e != nil {
		t.Failed()
	}

	// test read uint64 exceeds max
	r.Buffer = w.Buffer
	if e := r.SerializeUint(&i, MaxUint64-2); e == nil {
		t.FailNow()
	}

	// test read uint64
	r.Buffer = w.Buffer
	if e := r.SerializeUint(&i, MaxUint64); e != nil || i != o {
		t.FailNow()
	}

	// test write uint32
	w.Reset()
	o = uint(MaxUint32 - 1)
	if e := w.SerializeUint(&o, uint64(MaxUint32)); e != nil {
		t.FailNow()
	}

	// test read uint32 exceeds max
	r.Buffer = w.Buffer
	if e := r.SerializeUint(&i, uint64(MaxUint32-2)); e == nil {
		t.FailNow()
	}

	// test read uint32
	r.Buffer = w.Buffer
	if e := r.SerializeUint(&i, uint64(MaxUint32)); e != nil || i != o {
		t.FailNow()
	}

	// test write uint16
	w.Reset()
	o = uint(MaxUint16 - 1)
	if e := w.SerializeUint(&o, uint64(MaxUint16)); e != nil {
		t.FailNow()
	}

	// test read uint16 exceeds max
	r.Buffer = w.Buffer
	if e := r.SerializeUint(&i, uint64(MaxUint16-2)); e == nil {
		t.FailNow()
	}

	// test read uint16
	r.Buffer = w.Buffer
	if e := r.SerializeUint(&i, uint64(MaxUint16)); e != nil || i != o {
		t.FailNow()
	}

	// test write uint8
	w.Reset()
	o = uint(MaxUint8 - 1)
	if e := w.SerializeUint(&o, uint64(MaxUint8)); e != nil {
		t.FailNow()
	}

	// test read uint8 exceeds max
	r.Buffer = w.Buffer
	if e := r.SerializeUint(&i, uint64(MaxUint8-2)); e == nil {
		t.FailNow()
	}

	// test read uint8
	r.Buffer = w.Buffer
	if e := r.SerializeUint(&i, uint64(MaxUint8)); e != nil || i != o {
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
