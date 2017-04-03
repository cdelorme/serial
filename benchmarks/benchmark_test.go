package benchmarks

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"sync/atomic"
	"testing"

	"github.com/cdelorme/serial"
	"github.com/tinylib/msgp/msgp"
)

var minSerial, maxSerial, minGob, maxGob, minMsgp, maxMsgp atomic.Value

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestMain(m *testing.M) {
	maxSerial.Store(0)
	minSerial.Store(1000)
	maxGob.Store(0)
	minGob.Store(1000)
	minMsgp.Store(1000)
	maxMsgp.Store(0)
	exitCode := m.Run()
	fmt.Printf("serial sizes: %d - %d bytes\n", minSerial.Load(), maxSerial.Load())
	fmt.Printf("gob sizes:    %d - %d bytes\n", minGob.Load(), maxGob.Load())
	fmt.Printf("msgp sizes:   %d - %d bytes\n", minMsgp.Load(), maxMsgp.Load())
	os.Exit(exitCode)
}

func BenchmarkSerializePackets(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var packet bytes.Buffer
		d := NewEntity()
		benchReader := serial.NewReader(&packet)
		benchWriter := serial.NewWriter(&packet)

		if e := d.Serialize(benchWriter); e != nil {
			b.Errorf("failed to serialize write: %s\n", e)
		} else {
			minSerial.Store(min(minSerial.Load().(int), packet.Len()))
			maxSerial.Store(max(maxSerial.Load().(int), packet.Len()))
		}
		if e := d.Serialize(benchReader); e != nil {
			b.Errorf("failed to serialize read: %s\n", e)
		}
	}
}

func BenchmarkSerializeStream(b *testing.B) {
	b.ReportAllocs()
	var network bytes.Buffer
	benchReader := serial.NewReader(&network)
	benchWriter := serial.NewWriter(&network)

	for i := 0; i < b.N; i++ {
		d := NewEntity()
		if e := d.Serialize(benchWriter); e != nil {
			b.Errorf("failed to serialize write: %s\n", e)
		} else {
			minSerial.Store(min(minSerial.Load().(int), network.Len()))
			maxSerial.Store(max(maxSerial.Load().(int), network.Len()))
		}
		if e := d.Serialize(benchReader); e != nil {
			b.Errorf("failed to serialize read: %s\n", e)
		}
	}
}

func BenchmarkGobPackets(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var packet bytes.Buffer
		d := NewEntity()
		enc := gob.NewEncoder(&packet)
		dec := gob.NewDecoder(&packet)

		if e := enc.Encode(d); e != nil {
			b.Errorf("failed to gob encode message: %s\n", e)
		} else {
			minGob.Store(min(minGob.Load().(int), packet.Len()))
			maxGob.Store(max(maxGob.Load().(int), packet.Len()))
		}
		if e := dec.Decode(&d); e != nil {
			b.Errorf("failed to gob decode message: %s\n", e)
		}
	}
}

func BenchmarkGobStream(b *testing.B) {
	b.ReportAllocs()
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	for i := 0; i < b.N; i++ {
		d := NewEntity()
		if e := enc.Encode(d); e != nil {
			b.Errorf("failed to gob encode message: %s\n", e)
		} else {
			minGob.Store(min(minGob.Load().(int), network.Len()))
			maxGob.Store(max(maxGob.Load().(int), network.Len()))
		}
		if e := dec.Decode(&d); e != nil {
			b.Errorf("failed to gob decode message: %s\n", e)
		}
	}
}

func BenchmarkMsgpPacket(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var packet bytes.Buffer
		msgpReader := msgp.NewReader(&packet)
		msgpWriter := msgp.NewWriter(&packet)
		d := NewEntity()

		if e := d.EncodeMsg(msgpWriter); e != nil {
			b.Errorf("failed to msgp encode message: %s\n", e)
		} else {
			msgpWriter.Flush()
			minMsgp.Store(min(minMsgp.Load().(int), packet.Len()))
			maxMsgp.Store(max(maxMsgp.Load().(int), packet.Len()))
		}
		if e := d.DecodeMsg(msgpReader); e != nil {
			b.Errorf("failed to msgp decode message: %s\n", e)
		}
	}
}

func BenchmarkMsgpStream(b *testing.B) {
	b.ReportAllocs()
	var network bytes.Buffer
	msgpReader := msgp.NewReader(&network)
	msgpWriter := msgp.NewWriter(&network)

	for i := 0; i < b.N; i++ {
		d := NewEntity()
		if e := d.EncodeMsg(msgpWriter); e != nil {
			b.Errorf("failed to msgp encode message: %s\n", e)
		} else {
			msgpWriter.Flush()
			minMsgp.Store(min(minMsgp.Load().(int), network.Len()))
			maxMsgp.Store(max(maxMsgp.Load().(int), network.Len()))
		}
		if e := d.DecodeMsg(msgpReader); e != nil {
			b.Errorf("failed to msgp decode message: %s\n", e)
		}
	}
}
