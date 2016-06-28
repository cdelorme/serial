
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

This is my final performance update:

- explicitly define uint16 for each stat

My hope is that by using deterministic types I can improve performance of both the serialization and `encoding/gob` processes, while shrinking the size further for the `encoding/gob` implementation.  _Skipping past the `MaxSize` logic should shave some time off the processing giving us an optimized `Entity`._

The final benchmarks:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      1855 ns/op
	BenchmarkGob-8      	 1000000	      2262 ns/op
	ok  	github.com/cdelorme/go-udp-transport	4.173s

The final sizes:

	Serialized: 18
	Gobbed: 114

As I predicted, the size of the serialization did not change (since we optimized that with `MaxSize` logic previously), but the `encoding/gob` did shrink by a single byte, and the performance still remains over 400~ns faster for the serialization solution.

**In conclusion, I was able to write a proof-of-concept serialization tool for udp transport following the suggestions by Gaffer's articles, and his advice proved to be true, as just a few days of work is significantly faster and produced much more compact data.**


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
