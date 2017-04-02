package transport

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

const MaxSizeName uint64 = 255

type Serializer interface {
	SerializeString(*string, uint64) error
	SerializeBool(*bool) error
	SerializeUint16(*uint16) error
}

type Entity struct {
	Name    string
	Health  [2]uint16
	Mana    [2]uint16
	Stamina [2]uint16
	Dead    bool
}

func (o *Entity) Serialize(s Serializer) error {
	if e := s.SerializeString(&o.Name, MaxSizeName); e != nil {
		return e
	}
	if e := s.SerializeUint16(&o.Health[0]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&o.Health[1]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&o.Mana[0]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&o.Mana[1]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&o.Stamina[0]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&o.Stamina[1]); e != nil {
		return e
	}
	return s.SerializeBool(&o.Dead)
}

var benchEntity = Entity{
	Name:    "Casey",
	Health:  [2]uint16{100, 100},
	Mana:    [2]uint16{50, 50},
	Stamina: [2]uint16{75, 75},
}

func init() {
	var sn bytes.Buffer
	benchWriter := &WriteSerial{&sn}
	benchEntity.Serialize(benchWriter)
	fmt.Printf("Serialized Bytes: %d\n", benchWriter.Len())

	var en bytes.Buffer
	enc := gob.NewEncoder(&en)
	enc.Encode(benchEntity)
	fmt.Printf("Gobbed Bytes:     %d\n", len(en.Bytes()))
}

func TestEntity(t *testing.T) {
	t.Parallel()
	name := "Casey"
	var b bytes.Buffer
	o, i := Entity{Name: name, Health: [2]uint16{100, 1000}}, Entity{}
	r, w := &ReadSerial{Buffer: &bytes.Buffer{}}, &WriteSerial{Buffer: &b}

	// serialize data to write serial
	if e := o.Serialize(w); e != nil || w.Len() == 0 {
		t.FailNow()
	}

	// force error with invalid data
	r.Write(w.Bytes()[:8])
	if e := i.Serialize(r); e == nil {
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

	// de-serialize from read serial using previous write serial's data
	r.Buffer = w.Buffer
	if e := i.Serialize(r); e != nil || i.Name != name {
		t.FailNow()
	}
}

func BenchmarkSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var network bytes.Buffer
		benchReader := &ReadSerial{&network}
		benchWriter := &WriteSerial{&network}

		benchEntity.Serialize(benchWriter)
		benchEntity.Serialize(benchReader)
	}
}

func BenchmarkGobOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var network bytes.Buffer
		enc := gob.NewEncoder(&network)
		dec := gob.NewDecoder(&network)

		enc.Encode(benchEntity)
		dec.Decode(&benchEntity)
	}
}

func BenchmarkGobTwo(b *testing.B) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	// @note: this is not how the encoding system would be implemented
	// due to concurrency and multiple running systems that would deal
	// with messages in parallel so we didn't block the pipe, and we
	// would probably have to create a new "network" ([]byte) per request
	for i := 0; i < b.N; i++ {
		enc.Encode(benchEntity)
		dec.Decode(&benchEntity)
	}
}
