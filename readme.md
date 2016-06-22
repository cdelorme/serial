
# go-udp-transport

This was a test case to mirror the C++ implementation proposed by [Gaffer's Game Networking Protocol Articles](http://gafferongames.com/2016/05/10/building-a-game-network-protocol/).

These are the goals:

- demonstrate a minimalistic serialization implementation in go
- compare performance & efficiency of `encoding/gob` vs minimalistic serialization implementation
- identify deficiencies that could be surmounted with extra effort


## implementation

Due to the implicit interfaces, we can easily mirror the serialization strategy described by Gaffer.

This strategy relies on matching function parameters for a read and write utility.

For maximum efficiency we deal with pointers to avoid extra memory allocation overhead, while creating or reading from a byte array.


## performance

It is worth noting that one of the intended purposes of the `encoding/gob` package was efficient network traffic.  It is therefore pre-optimized.  Some important behaviors to note are that it ignores private struct fields, flattens pointers, and allocates space to define each data type.  Whether these are beneficial depends on the use-case, _but these are decisions made by the go team so I have to assume most are beneficial._

While my plan is to iterate for efficiency, the first draft will be naive array length storage, use `LittleEndian` for its `ByteOrder` with no byte-packing; it is likely to be less efficiency and possibly slower (although this is not for certain).  As I release each major update, I will update the benchmarks here:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkGob-8      	  500000	      2418 ns/op
	BenchmarkSerialize-8	 1000000	      2175 ns/op
	ok  	github.com/cdelorme/go-udp-transport	3.448s

_In complete opposition of my expectations, the gob library was not only slower than a custom serialization without any optimizations, but the bytes consumed were nearly half the gob solution (`61` vs `115`)._  I am a bit shocked since I expected a very different outcome.

**With a bit of cleanup and optimization, it's possible to make serialize way faster and better!**


## deficiencies

One major efficiency problem is that go's default size integer (eg. `int`) has a variable size that isn't compatible with the `encoding/binary` package.  The conversion process from types to bytes and back simply can't deal with that type.

Because that type is variable in size, for safety we have to assume int64.  This means the size of a byte array takes up 8 bytes before we start writing the array data.

There are three ways to deal with this problem:

- arbitrary fixed size restriction to int32 or int16 always
- add a function per size; eg. `SerializeString32`, `SerializeString16`, etc...
- accept a max size parameter and use that to efficiently handle length storage as well as errors for over-sized

_The same problem will be encountered when dealing with int properties, which means this problem extends quite a ways._  If we use the `MaxSize` approach, we have a nice clean way to deal with this problem and enforce validation at a much more sane level.

The next problem is organization of read and write stream in separate files doesn't line up well with matching tests since it's very difficult to test read and write independently, and we wouldn't want to if we're trying to find bugs with incompatibilities introduced (eg. changing one half's `ByteOrder`).


# references

- [gobs of data](https://blog.golang.org/gobs-of-data)
- [encoding/gob](https://golang.org/pkg/encoding/gob/)
- [encoding/binary](https://golang.org/pkg/encoding/binary/)
- [io](https://golang.org/pkg/io/)
- [go slices usage and internals](https://blog.golang.org/go-slices-usage-and-internals)
