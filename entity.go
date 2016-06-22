package transport

type Entity struct {
	Name    string
	Health  [2]int
	Mana    [2]int
	Stamina [2]int
}

func (self *Entity) Serialize(s Stream) error {
	if e := s.SerializeString(&self.Name); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Health[0]); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Health[1]); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Mana[0]); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Mana[1]); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Stamina[0]); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Stamina[1]); e != nil {
		return e
	}
	return nil
}
