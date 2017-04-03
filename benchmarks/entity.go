//go:generate msgp -tests=false
package benchmarks

type Entity struct {
	Name     string
	Health   [2]uint16
	Mana     [2]uint16
	Stamina  [2]uint16
	Friends  []int64
	Statuses []Status
	Dead     bool
}

func (o *Entity) Serialize(s Serializer) error {
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

	lf := int64(len(o.Friends))
	if e := s.Serialize(&lf); e != nil {
		return e
	}
	if int(lf) != len(o.Friends) {
		o.Friends = make([]int64, int(lf))
	}
	if e := s.Serialize(&o.Friends); e != nil {
		return e
	}

	ls := int64(len(o.Statuses))
	if e := s.Serialize(&ls); e != nil {
		return e
	}
	if int(ls) != len(o.Statuses) {
		o.Statuses = make([]Status, int(ls))
	}
	for i := range o.Statuses {
		o.Statuses[i].Serialize(s)
	}
	return s.Serialize(&o.Health, &o.Mana, &o.Stamina, &o.Dead)
}
