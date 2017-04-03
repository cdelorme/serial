
# [serial](https://github.com/cdelorme/serial)

This package is a proof-of-concept modeling a general purpose serialization strategy after the articles by [Glenn Fiedler](http://gafferongames.com/2016/05/10/building-a-game-network-protocol/).

It is intended to demonstrate the serialization pattern described, which aims to tightly pack bytes to minimize transport costs in a UDP communication system, while also providing a single function for both read and write to reduce potential for human error when modifications to the structure are made.

The code has seen several revisions, balancing execution performance, convenience, and the smallest byte packing possible while using the [`encoding/binary` package](https://golang.org/pkg/encoding/binary/) in `LittleEndian`.  The latest iteration has traded for convenience at the cost of reduced built-in functionality.  _This is mostly due to the lack of sane alternatives when dealing with variable size types, most especially complex combinations such as maps or unsized slices._

While it can achieve significantly smaller byte sizes than tools like [`encoding/gob`](https://golang.org/pkg/encoding/gob/) or even [`github.com/tinylib/msgp`](https://github.com/tinylib/msgp), it comes at the expense of spending many hours fine-tuning and writing the serialization logic by hand.

This code comes with a full set of unit tests, and the [benchmarks](benchmarks/) have been cleanly separated to create a more well defined example without mixing up code.


## conclusions

**With the MTU of only 1500 bytes, every byte counts when trying to create a communication protocol over UDP that does not fragment.**

_I have to agree with Glenn Fiedler regarding the fact that the best performance can only be achieved by explicitly defining serialization per structure._

However, depending on your use case the [msgp package](https://github.com/tinylib/msgp) offers absolutely stellar performance and very low memory consumption.  It produces results just under 2x the size of manual serialization, _but at zero cognitive overhead._  Using this lets you focus on the other parts of the application, and fits very nicely within the same overall packet-based strategy.


On the other hand, I cannot recommend the [`encoding/gob` package](https://golang.org/pkg/encoding/gob/).  It consumes significantly more memory, and sometimes upwards of 6x the amount of bytes due to the amount of metadata generated.  Additionally it is more geared towards streaming, and has heavy dependency on a per-instance cache.  _This model conflicts with packet based UDP, and new instances are an order of magnitude slower to parse or write messages._

To summarize, if you want to focus on the rest of your application just use the `msgp` package.  If you get to a point where you need to optimize, it takes very little effort to exchange it for custom serialization.  _However, writing custom serialization is very expensive._  It is best to plan according to your applications needs.


# references

- [gobs of data](https://blog.golang.org/gobs-of-data)
- [encoding/gob](https://golang.org/pkg/encoding/gob/)
- [encoding/binary](https://golang.org/pkg/encoding/binary/)
- [io](https://golang.org/pkg/io/)
- [go slices usage and internals](https://blog.golang.org/go-slices-usage-and-internals)
- [golang specification types](https://golang.org/ref/spec#Types)
