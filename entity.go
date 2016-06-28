package transport

const MaxSizeName uint64 = 255

type Entity struct {
	Name    string
	Health  [2]uint16
	Mana    [2]uint16
	Stamina [2]uint16
}

func (self *Entity) Serialize(s Stream) error {
	if e := s.SerializeString(&self.Name, MaxSizeName); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Health[0]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Health[1]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Mana[0]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Mana[1]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Stamina[0]); e != nil {
		return e
	}
	if e := s.SerializeUint16(&self.Stamina[1]); e != nil {
		return e
	}
	return nil
}
