package transport

import (
	"bytes"
	"testing"
)

func TestPlacebo(_ *testing.T) {}

func TestNewSerials(t *testing.T) {
	t.Parallel()
	if r := NewReadSerial(nil); r == nil {
		t.FailNow()
	}
	if w := NewWriteSerial(nil); w == nil {
		t.FailNow()
	}
}

func TestSerialString(t *testing.T) {
	t.Parallel()
	var o, i string = "bananas", ""
	var b bytes.Buffer
	r, w := &ReadSerial{&bytes.Buffer{}}, &WriteSerial{&b}

	// test decode empty data
	if e := r.SerializeString(&i, 0); e == nil {
		t.FailNow()
	}

	// test max length safety
	if e := w.SerializeString(&o, 2); e == nil {
		t.FailNow()
	}

	// test encoding (default)
	if e := w.SerializeString(&o, 0); e != nil || w.Len() == 0 {
		t.FailNow()
	}
	l := w.Len()

	// // test decode partial error
	r.Write(w.Bytes()[:8])
	if e := r.SerializeString(&i, 0); e == nil {
		t.FailNow()
	}

	// test full decode
	r.Buffer = w.Buffer
	if e := r.SerializeString(&i, 0); e != nil || o != i {
		t.FailNow()
	}

	// test write empty string
	w.Reset()
	var g string
	if e := w.SerializeString(&g, 0); e != nil || w.Len() == 0 {
		t.FailNow()
	}

	// test optimized encoding
	w.Reset()
	if e := w.SerializeString(&o, 255); e != nil || w.Len() >= l {
		t.FailNow()
	}
}

func TestSerialInt(t *testing.T) {
	t.Parallel()
	var o, i int
	var b bytes.Buffer
	r, w := ReadSerial{&bytes.Buffer{}}, WriteSerial{&b}

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
	// r.Buffer = w.Buffer
	r.Write(w.Bytes())
	if e := r.SerializeInt(&i, MaxInt64-2); e == nil {
		t.FailNow()
	}

	// fails here, shared buffer?
	// r.Buffer = w.Buffer
	r.Write(w.Bytes())
	if e := r.SerializeInt(&i, MaxInt64); e != nil || i != o {
		t.FailNow()
	}

	// test write int32
	w.Reset()
	o = int(MaxInt32 - 1)
	if e := w.SerializeInt(&o, int64(MaxInt32)); e != nil || w.Len() != 4 {
		t.FailNow()
	}

	// test read int32 exceeds max
	r.Write(w.Bytes())
	if e := r.SerializeInt(&i, int64(MaxInt32-2)); e == nil {
		t.FailNow()
	}

	// test read int32
	r.Write(w.Bytes())
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
	r.Write(w.Bytes())
	if e := r.SerializeInt(&i, int64(MaxInt16-2)); e == nil {
		t.FailNow()
	}

	// test read int16
	r.Write(w.Bytes())
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
	r.Write(w.Bytes())
	if e := r.SerializeInt(&i, int64(MaxInt8-2)); e == nil {
		t.FailNow()
	}

	// test read int8
	r.Write(w.Bytes())
	if e := r.SerializeInt(&i, int64(MaxInt8)); e != nil {
		t.FailNow()
	}
}

func TestSerialUint(t *testing.T) {
	t.Parallel()
	var o, i uint
	var b bytes.Buffer
	r, w := ReadSerial{&bytes.Buffer{}}, WriteSerial{&b}

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
	r.Write(w.Bytes())
	if e := r.SerializeUint(&i, MaxUint64-2); e == nil {
		t.FailNow()
	}

	// test read uint64
	r.Write(w.Bytes())
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
	r.Write(w.Bytes())
	if e := r.SerializeUint(&i, uint64(MaxUint32-2)); e == nil {
		t.FailNow()
	}

	// test read uint32
	r.Write(w.Bytes())
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
	r.Write(w.Bytes())
	if e := r.SerializeUint(&i, uint64(MaxUint16-2)); e == nil {
		t.FailNow()
	}

	// test read uint16
	r.Write(w.Bytes())
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
	r.Write(w.Bytes())
	if e := r.SerializeUint(&i, uint64(MaxUint8-2)); e == nil {
		t.FailNow()
	}

	// test read uint8
	r.Write(w.Bytes())
	if e := r.SerializeUint(&i, uint64(MaxUint8)); e != nil || i != o {
		t.FailNow()
	}
}

func TestSerialInt8(t *testing.T) {
	t.Parallel()
	var o, i int8 = 12, 0
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}

	if e := r.SerializeInt8(&i); e == nil {
		t.FailNow()
	}

	if e := w.SerializeInt8(&o); e != nil || w.Len() != 1 {
		t.FailNow()
	}

	if e := r.SerializeInt8(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestSerialInt16(t *testing.T) {
	t.Parallel()
	var o, i int16 = 12, 0
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}

	if e := r.SerializeInt16(&i); e == nil {
		t.FailNow()
	}

	if e := w.SerializeInt16(&o); e != nil || w.Len() != 2 {
		t.FailNow()
	}

	if e := r.SerializeInt16(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestSerialInt32(t *testing.T) {
	t.Parallel()
	var o, i int32 = 12, 0
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}

	if e := r.SerializeInt32(&i); e == nil {
		t.FailNow()
	}

	if e := w.SerializeInt32(&o); e != nil || w.Len() != 4 {
		t.FailNow()
	}

	if e := r.SerializeInt32(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestSerialInt64(t *testing.T) {
	t.Parallel()
	var o, i int64 = 12, 0
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}

	if e := r.SerializeInt64(&i); e == nil {
		t.FailNow()
	}

	if e := w.SerializeInt64(&o); e != nil || w.Len() != 8 {
		t.FailNow()
	}

	if e := r.SerializeInt64(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestSerialUint8(t *testing.T) {
	t.Parallel()
	var o, i uint8 = 12, 0
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}

	if e := r.SerializeUint8(&i); e == nil {
		t.FailNow()
	}

	if e := w.SerializeUint8(&o); e != nil || w.Len() != 1 {
		t.FailNow()
	}

	if e := r.SerializeUint8(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestSerialUint16(t *testing.T) {
	t.Parallel()
	var o, i uint16 = 12, 0
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}

	if e := r.SerializeUint16(&i); e == nil {
		t.FailNow()
	}

	if e := w.SerializeUint16(&o); e != nil || w.Len() != 2 {
		t.FailNow()
	}

	if e := r.SerializeUint16(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestSerialUint32(t *testing.T) {
	t.Parallel()
	var o, i uint32 = 12, 0
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}

	if e := r.SerializeUint32(&i); e == nil {
		t.FailNow()
	}

	if e := w.SerializeUint32(&o); e != nil || w.Len() != 4 {
		t.FailNow()
	}

	if e := r.SerializeUint32(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestSerialUint64(t *testing.T) {
	t.Parallel()
	var o, i uint64 = 12, 0
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}

	if e := r.SerializeUint64(&i); e == nil {
		t.FailNow()
	}

	if e := w.SerializeUint64(&o); e != nil || w.Len() != 8 {
		t.FailNow()
	}

	if e := r.SerializeUint64(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestSerialFloat32(t *testing.T) {
	t.Parallel()
	var o, i float32 = 12.3, 0
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}

	if e := r.SerializeFloat32(&i); e == nil {
		t.FailNow()
	}

	if e := w.SerializeFloat32(&o); e != nil || w.Len() != 4 {
		t.FailNow()
	}

	if e := r.SerializeFloat32(&i); e != nil || i != o {
		t.FailNow()
	}
}

func TestSerialFloat64(t *testing.T) {
	t.Parallel()
	var o, i float64 = 12.3, 0
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}

	if e := r.SerializeFloat64(&i); e == nil {
		t.FailNow()
	}

	if e := w.SerializeFloat64(&o); e != nil || w.Len() != 8 {
		t.FailNow()
	}

	if e := r.SerializeFloat64(&i); e != nil || i != o {
		t.FailNow()
	}
}
