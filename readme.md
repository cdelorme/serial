
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

This time I tried a few bigger changes:

- attempting to reuse existing `Read()`, `Write()`, and `SerializeInt()`
- switching to `bytes.Buffer` from custom

Attempting to re-use `Read()` very nearly added another 100ns back to the metrics, while a call to `Write()` was barely 5~ns added.  _While code reuse is appealing, this proves that it has a cost and my objective is efficiency so I'd like to find a way to avoid incurring that cost._

I decided to try and use the `bytes.Buffer` instead of a custom `Read()`, `Write()` and byte array with position handler.  Not only does this reduce code, it replaces it with existing reliable built-in code.  _It also got us the desired code reuse without the performance hit we had from before._

The resulting benchmarks with 18 less lines of code:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      1950 ns/op
	BenchmarkGob-8      	  500000	      2385 ns/op
	ok  	github.com/cdelorme/go-udp-transport	4.191s

_The benchmarks may not do justice to the serialization performance since I had extra steps to copy the data from the writer to the reader per operation, which theoretically would not be incurred in normal flow.  The `gob` library uses a pointer which yields a shared buffer, but the pointer dereferencing may be what reduces the `gob` performance._


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
- [golang specification types](https://golang.org/ref/spec#Types)
