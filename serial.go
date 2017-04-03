// This package is a proof-of-concept demonstrating a serialization pattern
// as described by Glenn Fiedlers articles on UDP network traffic.
//
// It has gone through several iterations to try new things, both to improve
// the performance, as well as the byte-packing.
//
// This latest iteration discards variable support in favor of simplification.
//
// The benchmarks have been enhanced to randomly populate a complex structure
// and use both tinylib/msgp and encoding/gob to compare both performance and
// byte size.
//
// While I have strategies to running serialization against a variety of data
// I have not yet identified a sane pattern for complex variable size types,
// such as maps, and slices of variable length or that contain variable length
// data.
//
// The benchmarks show significant packing when using this serialization tool,
// which can be further optimized through careful consideration when writing
// the serialization logic per structure.  However, this process is cognitively
// expensive.
//
// I cannot recommend the gob package.  It has enormous variable sizes due to
// the amount of metadata stored, and in scenarios where you cannot rely on an
// existing instance it performs quite horribly by depending on a cache per
// instance, without which significantly more effort (likely via reflect) and
// more allocations are required.
//
// The msgp solution is highly recommended for many reasons.  First it performs
// far better with zero effort than manual serialization.  Anywhere from 30% to
// several times faster.  It also has roughly 5 times less allocations, which
// is much lighter on memory consumption.  It supports a very similar pattern
// as the serialization strategy, and would be trivial to connect to a gzip
// package for example.  The first downside is that it consumes around twice
// as many bytes as serialization.  The second is that it produces a sizable
// amount of generated code.  The trade-off is that generated code takes zero
// effort away from the core development.  At worst, msgp is the absolute best
// way to get started with a project that needs good network communication.
package serial

import "encoding/binary"

var byteOrder = binary.LittleEndian

// A replica of the io.Writer interface for compatibility minus the import.
type Writer interface {
	Write([]byte) (int, error)
}

// A replica of the io.Reader interface for compatibility minus the import.
type Reader interface {
	Read([]byte) (int, error)
}

// Returns a new serialization reader with functionality matching the writer.
func NewReader(r Reader) *Read {
	return &Read{r}
}

// Returns a new serialization writer with functionality matching the reader.
func NewWriter(w Writer) *Write {
	return &Write{w}
}
