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

func (self *Entity) Serialize(s Serializer) error {
	if e := s.SerializeString(&self.Name, MaxSizeName); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Health[0]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Health[1]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Mana[0]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Mana[1]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Stamina[0]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Stamina[1]); e != nil {
		return e
	}
	if e := s.SerializeBool(&self.Dead); e != nil {
		return e
	}
	return nil
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

func BenchmarkGob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var network bytes.Buffer
		enc := gob.NewEncoder(&network)
		dec := gob.NewDecoder(&network)

		enc.Encode(benchEntity)
		dec.Decode(&benchEntity)
	}
}
