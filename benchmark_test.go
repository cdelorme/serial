package transport

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

var benchEntity = Entity{
	Name:    "Casey",
	Health:  [2]uint16{100, 100},
	Mana:    [2]uint16{50, 50},
	Stamina: [2]uint16{75, 75},
}

func init() {
	var sn bytes.Buffer
	benchWriter := &WriteStream{&sn}
	benchEntity.Serialize(benchWriter)
	fmt.Printf("Serialized Bytes: %d\n", benchWriter.Len())

	var en bytes.Buffer
	enc := gob.NewEncoder(&en)
	enc.Encode(benchEntity)
	fmt.Printf("Gobbed Bytes:     %d\n", len(en.Bytes()))
}

func BenchmarkSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var network bytes.Buffer
		benchReader := &ReadStream{&network}
		benchWriter := &WriteStream{&network}

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
