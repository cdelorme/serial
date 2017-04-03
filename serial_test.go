package serial

import (
	"bytes"
	"testing"
)

func TestPlacebo(_ *testing.T) {}

func TestNewSerials(t *testing.T) {
	t.Parallel()
	if r := NewReader(nil); r == nil {
		t.FailNow()
	}
	if w := NewWriter(nil); w == nil {
		t.FailNow()
	}
}

func TestSerialize(t *testing.T) {

	// supported types
	var o8, i8 int8 = -2, 0
	var o16, i16 int16 = -2, 0
	var o32, i32 int32 = -2, 0
	var o64, i64 int64 = -2, 0
	var uo8, ui8 uint8 = 2, 0
	var uo16, ui16 uint16 = 2, 0
	var uo32, ui32 uint32 = 2, 0
	var uo64, ui64 uint64 = 2, 0
	var fo32, fi32 float32 = 0.12, 0
	var fo64, fi64 float64 = 0.12, 0
	var co64, ci64 complex128 = 0.12, 0
	var co128, ci128 complex128 = 0.12, 0
	var ro, ri rune = 'o', 'i'

	// shared buffer and serializers
	var b bytes.Buffer
	r, w := &Read{&b}, &Write{&b}

	// ignore emtpy input
	if e := w.Serialize(); e != nil {
		t.Error("failed to ignore empty input on write...")
	}
	if e := r.Serialize(); e != nil {
		t.Error("failed to ignore empty input on read...")
	}

	// all successful cases in one shot
	if e := w.Serialize(
		&o8,
		&o16,
		&o32,
		&o64,
		&uo8,
		&uo16,
		&uo32,
		&uo64,
		&fo32,
		&fo64,
		&co64,
		&co128,
		&ro,
	); e != nil {
		t.Errorf("failed to write: %s\n", e)
	} else if b.Len() == 0 {
		t.Error("failed to write bytes and no error was received...")
	}
	if e := r.Serialize(
		&i8,
		&i16,
		&i32,
		&i64,
		&ui8,
		&ui16,
		&ui32,
		&ui64,
		&fi32,
		&fi64,
		&ci64,
		&ci128,
		&ri,
	); e != nil {
		t.Errorf("failed to read: %s\n", e)
	} else {
		if i8 != o8 {
			t.Errorf("failed to restore %T value: %v, expected %v", i8, i8, o8)
		}
		if i16 != o16 {
			t.Errorf("failed to restore %T value: %v, expected %v", i16, i16, o16)
		}
		if i32 != o32 {
			t.Errorf("failed to restore %T value: %v, expected %v", i32, i32, o32)
		}
		if i64 != o64 {
			t.Errorf("failed to restore %T value: %v, expected %v", i64, i64, o64)
		}
		if ui8 != uo8 {
			t.Errorf("failed to restore %T value: %v, expected %v", ui8, ui8, uo8)
		}
		if ui16 != uo16 {
			t.Errorf("failed to restore %T value: %v, expected %v", ui16, ui16, uo16)
		}
		if ui32 != uo32 {
			t.Errorf("failed to restore %T value: %v, expected %v", ui32, ui32, uo32)
		}
		if ui64 != uo64 {
			t.Errorf("failed to restore %T value: %v, expected %v", ui64, ui64, uo64)
		}
		if fi32 != fo32 {
			t.Errorf("failed to restore %T value: %v, expected %v", fi32, fi32, fo32)
		}
		if fi64 != fo64 {
			t.Errorf("failed to restore %T value: %v, expected %v", fi64, fi64, fo64)
		}
		if ci64 != co64 {
			t.Errorf("failed to restore %T value: %v, expected %v", ci64, ci64, co64)
		}
		if ci128 != co128 {
			t.Errorf("failed to restore %T value: %v, expected %v", ci128, ci128, co128)
		}
		if ri != ro {
			t.Errorf("failed to restore %T value: %v, expected %v", ri, ri, ro)
		}
	}

	// failure cases
	var badString string
	var badInt int
	var badUint uint
	if e := r.Serialize(&badString); e == nil {
		t.Error("failed to empty read error...")
	}
	if e := w.Serialize(&badString); e == nil {
		t.Error("failed to capture error when writing unfixed data type %T\n", badString)
	}
	if e := w.Serialize(&badInt); e == nil {
		t.Error("failed to capture error when writing unfixed data type %T\n", badInt)
	}
	if e := w.Serialize(&badUint); e == nil {
		t.Error("failed to capture error when writing unfixed data type %T\n", badUint)
	}
}
