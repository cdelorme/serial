
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

I tried a few things:

- making `ByteOrder` a package variable and comparing `BigEndian` to `LittleEndian`
- optimizing the `Read()` to grab the whole byte chunk by parameter length

While this may not be the case over the network, the byte order made no difference to the size nor the performance.  _Still 61 bytes anywhere from 2200~2240ns regardless of ByteOrder used._  The go network stack should automatically convert the byte order to `BigEndian`, so this really wouldn't matter unless I bypassed the entire default `net` package.

Optimizing the `Read` to use the length of the slice to read the bytes instead of one-byte-at-a-time was a huge win, dropping nearly 200~ns for the serialization approach.

New Benchmarks:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      1911 ns/op
	BenchmarkGob-8      	  500000	      2650 ns/op
	ok  	github.com/cdelorme/go-udp-transport	4.289s


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
