
# go-udp-transport

This was a test case to mirror the C++ implementation proposed by [Gaffer's Game Networking Protocol Articles](http://gafferongames.com/2016/05/10/building-a-game-network-protocol/).

These are the goals:

- demonstrate a minimalistic serialization implementation in go
- compare performance & efficiency of `encoding/gob` vs minimalistic serialization implementation
- identify problems that could be surmounted with extra effort


## implementation

Due to the implicit interfaces, we can easily mirror the serialization strategy described by Gaffer.

This strategy relies on matching function parameters for a read and write utility.

For maximum efficiency we deal with pointers to avoid extra memory allocation overhead, while creating or reading from a byte array.  _The process of converting data into byte arrays for writing into a buffer requires allocations, so we aren't saving a whole lot by using pointers for writes._


## performance

This update I decided to make some changes to optimize usage and not performance or compression:

- use `*bytes.Buffer` and add `NewReadStream([]byte)` & `NewWriteStream([]byte)` for simple creation
- use `init()` in `benchmark_test.go` to print sizes instead of a test function
- modifying benchmarks to accurately reflect implementation

This modification makes it much easier to create and share state between copies of read and write streams, especially for the UDP use-case where we get a byte array and would like to directly use it when creating a `bytes.Buffer` without consuming more space.  _Perhaps due to less load on my system overall both benchmarks increased in performance, so it appears to have had little to no impact on the streams._

The second modification had a surprising impact, boosting benchmark performance of the `encoding/gob` package by nearly 300ns per operation.  _I assume this is related to the fact that the first run set the size of the `bytes.Buffer` resulting in better performance, but I'm not entirely sure._

To support a highly concurrent network application we would want a separate encoder/decoder per connection or component, or perhaps even per request.  To avoid large allocations of the same space we would want to directly set a slice of bytes from a channel or mutex-protected `ReadFromUDP()` onto a new `bytes.Buffer` or similar `ReadWriter`.  Clients would distribute messages by system component, and servers would distribute messages by connected client then system component.  _Somehow creating a new `encoding/gob` encoder and decoder per operation results in a dramatic reduction in performance._

Benchmarks with `*bytes.Buffer`:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      1874 ns/op
	BenchmarkGob-8      	 1000000	      2263 ns/op
	ok  	github.com/cdelorme/go-udp-transport	4.194s

Benchmarks with `init()`:

	$ go test -v -run=X -bench=.
	Serialized: 18
	Gobbed: 114
	PASS
	BenchmarkSerialize-8	 1000000	      1861 ns/op
	BenchmarkGob-8      	 1000000	      1987 ns/op
	ok  	github.com/cdelorme/go-udp-transport	3.899s

Benchmarks with new variables per loop:

	$ go test -v -run=X -bench=.
	Serialized Bytes: 18
	Gobbed Bytes:     114
	PASS
	BenchmarkSerialize-8	 1000000	      1985 ns/op
	BenchmarkGob-8      	   30000	     40104 ns/op
	ok  	github.com/cdelorme/go-udp-transport	3.632s

**Since UDP is not intended to be a stream and because we cannot capture UDP header address information, we cannot directly connect the the streams to the UDP connection, nor the encoding tools.**  The results above further prove that for a high performance application a serialization model that can be initialized in a distributed model with minimal changes to performance is ideal.


## problems

One significant problem is that the `int` data type, a common default in go, does not have a deterministic size.  The result is that it cannot be directly encoded using the `encoding/binary` package.  For safety you have to assume the largest type (int64) and add extra casting logic when dealing with it.  _This becomes a brand new set of problems when comparing results from the build-in `len()` function._

One particular case to worry about is getting the `len()` of a string after converting it to a byte array, so that we can correctly read the bytes back in from the `ReadStream` entity.  We have to not only assume such a large size, but now we have to store a minimum of 8 bytes for that size.

There are a few ways to deal with this problem:

- arbitrary fixed size restriction to int32 or int16 always
- add a function per size; eg. `SerializeString32`, `SerializeString16`, etc...
- accept a max size parameter and use that to efficiently handle length storage as well as errors for over-sized

My solution is to combine the creation of explicit sized integer storage with a maximum size parameter on non-deterministic types such as `string` or `int`.  If a zero value is supplied we assume the maximum, else we work within the boundaries of the supplied size.

**This gives us the benefit of more efficient storage at minimal cost of extra logic, as well as the ability to preemptively filter invalid values on behalf of the user.**

One of the other problems is that the organization of read and write streams doesn't align well with tests, since it's difficult to independently test read from write unless you know how to manually create byte arrays with valid integer and string data.  Additionally attempting to find bugs created by desynchronization between both constructs would be more difficult if tests did not use both entities, such as difference of `ByteOrder` or sizes used.


# references

- [gobs of data](https://blog.golang.org/gobs-of-data)
- [encoding/gob](https://golang.org/pkg/encoding/gob/)
- [encoding/binary](https://golang.org/pkg/encoding/binary/)
- [io](https://golang.org/pkg/io/)
- [go slices usage and internals](https://blog.golang.org/go-slices-usage-and-internals)
- [golang specification types](https://golang.org/ref/spec#Types)
