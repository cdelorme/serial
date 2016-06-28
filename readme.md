
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

I added a significant amount of code this time to try and optimize the space consumed:

- added functions for deterministic types: float32/64, int8/16/32/64, uint8/16/32/64
- added `MaxSize` parameter to `SerializeString()`, `SerializeInt()`, and `SerializeUint()`

Since I did not (_yet_) modify the entity encoding I wasn't expecting a change to performance, but I ended up seeing a cost of around 10~30ns per operation.  _Fortunately adding code reuse did not appear to create any difference in performance._

By adding the new `MaxSize` parameter I expected to take a significant hit to performance, but also to achieve far better efficiency when storing data.  In addition, the parameter can be used to invoke a user-defined maximum accepted value, reducing extra load post-serialization.  _To my surprise the performance only jumped between 50~100ns per operation, still remaining around 400~ns faster than the `encoding/gob` implementation, but the size has now dropped by 75% bringing us to 1/6th the size of the `encoding/gob` results._

Benchmarks with `MaxSize` modifications:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      2089 ns/op
	BenchmarkGob-8      	  500000	      2535 ns/op
	ok  	github.com/cdelorme/go-udp-transport	3.417s

New byte sizes per encoding method:

	Serialized: 18
	Gobbed: 115

_This time we added a total of `126` lines; `307` if we count tests, possibly less it we omit the entity that exists for testing the library._


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
