package transport

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type Serializer interface {
	Serialize(...interface{}) error
}

type Entity struct {
	Name    string
	Health  [2]uint16
	Mana    [2]uint16
	Stamina [2]uint16
	Dead    bool
}

func (o *Entity) Serialize(s Serializer) error {
	l := int8(len(o.Name))
	if e := s.Serialize(&l); e != nil {
		return e
	}
	b := []byte(o.Name)
	if int(l) != len(o.Name) {
		b = make([]byte, l)
	}
	if e := s.Serialize(&b); e != nil {
		return e
	}
	o.Name = string(b)
	return s.Serialize(&o.Health, &o.Mana, &o.Stamina, &o.Dead)
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
	fmt.Printf("Serialized Bytes: %d\n", sn.Len())

	var en bytes.Buffer
	enc := gob.NewEncoder(&en)
	enc.Encode(benchEntity)
	fmt.Printf("Gobbed Bytes:     %d\n", en.Len())
}

func TestEntity(t *testing.T) {
	var b bytes.Buffer
	r, w := &ReadSerial{&b}, &WriteSerial{&b}
	o, i := Entity{Name: "Casey", Health: [2]uint16{100, 1000}}, Entity{}

	if e := i.Serialize(r); e == nil {
		t.Error("failed to capture error with no data...")
	}
	if e := o.Serialize(w); e != nil || b.Len() == 0 {
		if e != nil {
			t.Logf("Error: %s\n", e)
		}
		t.Errorf("failed to write structure (%T)...\n", o)
	}
	if e := i.Serialize(r); e != nil || o.Name != i.Name || o.Health[0] != i.Health[0] || o.Health[1] != i.Health[1] {
		t.Errorf("failed to read structure (%T)...\n", o)
	}
	var d bytes.Buffer
	r.r = &d
	for n := 0; n < b.Len(); n++ {
		d.Reset()
		d.Write(b.Bytes()[:n])
		if e := i.Serialize(r); e == nil {
			t.Error("failed to receive error on partial entity data...")
			break
		}
	}
}

func BenchmarkSerializePackets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var packet bytes.Buffer
		benchReader := &ReadSerial{&packet}
		benchWriter := &WriteSerial{&packet}

		benchEntity.Serialize(benchWriter)
		benchEntity.Serialize(benchReader)
	}
}

func BenchmarkSerializeStream(b *testing.B) {
	var network bytes.Buffer
	benchReader := &ReadSerial{&network}
	benchWriter := &WriteSerial{&network}

	for i := 0; i < b.N; i++ {
		benchEntity.Serialize(benchWriter)
		benchEntity.Serialize(benchReader)
	}
}

func BenchmarkGobPackets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var packet bytes.Buffer
		enc := gob.NewEncoder(&packet)
		dec := gob.NewDecoder(&packet)

		enc.Encode(benchEntity)
		dec.Decode(&benchEntity)
	}
}

func BenchmarkGobStream(b *testing.B) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	for i := 0; i < b.N; i++ {
		enc.Encode(benchEntity)
		dec.Decode(&benchEntity)
	}
}
