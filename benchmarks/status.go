//go:generate msgp -tests=false
package benchmarks

type Status struct {
	Name     string
	Duration uint16
}

func (o *Status) Serialize(s Serializer) error {
	ln := int8(len(o.Name))
	if e := s.Serialize(&ln); e != nil {
		return e
	}
	b := []byte(o.Name)
	if int(ln) != len(o.Name) {
		b = make([]byte, int(ln))
	}
	if e := s.Serialize(&b); e != nil {
		return e
	}
	o.Name = string(b)
	return s.Serialize(&o.Duration)
}
