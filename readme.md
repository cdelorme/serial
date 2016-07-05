
# go-udp-transport

This is a library modeled off of a C++ implementation proposed by [Glenn Fiedler's Game Networking Protocol Articles](http://gafferongames.com/2016/05/10/building-a-game-network-protocol/), for efficient data serialization to send over a network using UDP.


## implementation

With implicit interfaces in go it is trivial to create two entities that match a single interface for serialization, and perform opposing operations.

To maximize efficiency we will allow for maximum sizes to be supplied when serializing a generic string (eg. a byte array), or generic `int` or `uint` data type.  _This allows us to use the smallest storage size for each type and enforce size checks on read to validate user-defined limits._

Finally we make use of pointers when dealing with all data to reduce allocations in as many places as possible.


### problems

During the implementation we ran into problems with `int` and `uint` types, forcing assumption of the largest data type to avoid loss (eg. `int64` and `uint64`).  This is because different architectures end up with varying sizes for these types, and the `encoding/binary` package refused to directly encode/decode them.  We were able to address storage efficiency by asking for a `MaxSize` when dealing with these generic types.

I ended up using `LittleEndian` for the `ByteOrder`, even though network transmission probably switches to `BigEndian`.  _I did try `BigEndian` and it made no noticeable difference in performance from benchmarks._


## design

My initial design took a naive approach of assuming the largest fixed data sizes (eg. `int64` and `uint64`), and automatically cast types during serialization.

For comparison I used the `encoding/gob` package, which initially had very similar performance but used more space, except when implemented in the same way that we expect a UDP system to work (where there is a mechanism per goroutine).

Here are just some of the steps I took to improve the initial design:

- making `ByteOrder` a package variable and comparing `BigEndian` to `LittleEndian`
- optimizing the `Read()` to grab the whole byte chunk by parameter length
- attempting to reuse existing `Read()`, `Write()`, and `SerializeInt()`
- switching to `bytes.Buffer` from custom
- added functions for deterministic types: float32/64, int8/16/32/64, uint8/16/32/64
- added `MaxSize` parameter to `SerializeString()`, `SerializeInt()`, and `SerializeUint()`
- explicitly define uint16 for each stat in the `Entity` struct
- use `*bytes.Buffer` and add `NewReadStream([]byte)` & `NewWriteStream([]byte)` for simple creation
- use `init()` in `benchmark_test.go` to print sizes instead of a test function
- modifying benchmarks to accurately reflect implementation

_The end result is a fully tested and very usable serialization library that takes up a fraction of the space consumed by `encoding/gob`, and fits very nicely into a UDP messaging model, but requires a bit more code to implement._


### performance

The first-draft benchmarks:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      2175 ns/op
	BenchmarkGob-8      	  500000	      2418 ns/op
	ok  	github.com/cdelorme/go-udp-transport	3.448s

_This first draft had a size of `115` bytes for `encoding/gob` and `61` bytes when using serialization._

After setting the `ByteOrder` via a package variable, and optimize the `Read()` behavior to do more than one byte at a time:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      1911 ns/op
	BenchmarkGob-8      	  500000	      2650 ns/op
	ok  	github.com/cdelorme/go-udp-transport	4.289s

Switching to a `bytes.Buffer` from custom `Read()` and `Write()` logic and implementing better code reuse:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      1950 ns/op
	BenchmarkGob-8      	  500000	      2385 ns/op
	ok  	github.com/cdelorme/go-udp-transport	4.191s

Adding functions for varying sizes of `int` and `uint`, plus optimizing storage with `MaxSize` settings:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      2089 ns/op
	BenchmarkGob-8      	  500000	      2535 ns/op
	ok  	github.com/cdelorme/go-udp-transport	3.417s

_At this point we dropped the serialization size from `61` to `18` bytes, making it literally smaller than 1/6th of the `encoding/gob` size._

Next I tried optimizing the `Entity` in the test case to use an `int16` directly to reduce logic and expected storage:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      1855 ns/op
	BenchmarkGob-8      	 1000000	      2262 ns/op
	ok  	github.com/cdelorme/go-udp-transport	4.173s

_This dropped the `encoding/gob` package size from `115` bytes to `114` bytes, while our existing serialization size remained at `18` bytes with a significant improvement in speed._

After wwitching to `*bytes.Buffer` to make the system more flexible to the expected UDP model:

	$ go test -v -run=X -bench=.
	PASS
	BenchmarkSerialize-8	 1000000	      1874 ns/op
	BenchmarkGob-8      	 1000000	      2263 ns/op
	ok  	github.com/cdelorme/go-udp-transport	4.194s

Somehow, moving the sizes to `init()` inside the benchmarks pre-sized the buffer dramatically improving the `encoding/gob` performance:

	$ go test -v -run=X -bench=.
	Serialized: 18
	Gobbed: 114
	PASS
	BenchmarkSerialize-8	 1000000	      1861 ns/op
	BenchmarkGob-8      	 1000000	      1987 ns/op
	ok  	github.com/cdelorme/go-udp-transport	3.899s

This final test reflects an actual expected use case where either new connections at the server or distributed components in a system would either have a brand new serialization instance or encoding tool, which absolutely destroyed our `encoding/gob` performance while having almost no effect on the serialization implementation:

	$ go test -v -run=X -bench=.
	Serialized Bytes: 18
	Gobbed Bytes:     114
	PASS
	BenchmarkSerialize-8	 1000000	      1985 ns/op
	BenchmarkGob-8      	   30000	     40104 ns/op
	ok  	github.com/cdelorme/go-udp-transport	3.632s

**I wanted to point out that the gob package is efficient if implemented correctly,** but that approach conflicts with a UDP pattern where each message is independent and not part of a single stream, and messages would probably be processed concurrently by various parts of the system:

	$ go test -v -run=X -bench=.
	Serialized Bytes: 19
	Gobbed Bytes:     123
	PASS
	BenchmarkSerialize-8	  500000	      2311 ns/op
	BenchmarkGobOne-8   	   30000	     45699 ns/op
	BenchmarkGobTwo-8   	 1000000	      2287 ns/op
	ok  	github.com/cdelorme/go-udp-transport	5.332s

_In this case we've added a boolean value to our entity, which has increased the `gob/encoding` size while resulting in better performance from that package._  The `gob/encoding` package should not be entirely ruled out, although it may also be faster to run serialization of booleans without code reuse (but I don't plan on tinkering at this point).


## conclusion

Even if we overlook the implementation details which effect the performance prediction (eg. as reflected in the final benchmarks), we were still able to create a serialization package with minimal effort that greatly exceeded both the alternative and expectation.

If we consider the network impact when shrinking from `114` to `18` bytes per object in a message in addition to the best case benchmarks of either solution, it's obvious that the winning scenario leans towards a custom serialization tool.

_However, there are still plenty of edge-cases that I have not yet covered with this library._

My conclusion is that Glenn Fiedler's suggestions and implementation are not only theoretically true but proven true from the very first iteration of this project.  _It may also be worth adding that the first iteration took only 3 hours, and this near-complete iteration roughly a week._


## future

One such edge-case to consider solving for is embedding byte arrays or arrays of structs in a way that is able to avoid desynchronization.  The current solution would require extra logic to optionally set or load the size of the array being serialized, and for simply storing an existing arbitrary byte array we would have to convert to a string first (which is a bit fugly).

I may also play with optimizing specific functionality without code reuse to try to speed up the code.


# references

- [gobs of data](https://blog.golang.org/gobs-of-data)
- [encoding/gob](https://golang.org/pkg/encoding/gob/)
- [encoding/binary](https://golang.org/pkg/encoding/binary/)
- [io](https://golang.org/pkg/io/)
- [go slices usage and internals](https://blog.golang.org/go-slices-usage-and-internals)
- [golang specification types](https://golang.org/ref/spec#Types)
