package transport

import (
	"bytes"
	"encoding/gob"
	"testing"
)

var benchEntity = Entity{
	Name:    "Casey",
	Health:  [2]int{100, 100},
	Mana:    [2]int{50, 50},
	Stamina: [2]int{75, 75},
}

var benchReader = &ReadStream{}
var benchWriter = &WriteStream{}

var network bytes.Buffer
var enc = gob.NewEncoder(&network)
var dec = gob.NewDecoder(&network)

func TestSizes(t *testing.T) {
	benchEntity.Serialize(benchWriter)
	t.Logf("Serialized: %d\n", benchWriter.Len())
	benchWriter.Reset()

	enc.Encode(benchEntity)
	t.Logf("Gobbed: %d\n", len(network.Bytes()))
	network.Reset()
}

func BenchmarkSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchEntity.Serialize(benchWriter)
		benchReader.Write(benchWriter.Bytes())
		benchEntity.Serialize(benchReader)
		benchWriter.Reset()
		benchReader.Reset()
	}
}

func BenchmarkGob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		enc.Encode(benchEntity)
		dec.Decode(&benchEntity)
	}
}
