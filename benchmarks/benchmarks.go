package benchmarks

type Serializer interface {
	Serialize(...interface{}) error
}
