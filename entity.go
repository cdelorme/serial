package transport

const MaxSizeName uint64 = 255
const MaxSizeStats int64 = 10000

type Entity struct {
	Name    string
	Health  [2]int
	Mana    [2]int
	Stamina [2]int
}

func (self *Entity) Serialize(s Stream) error {
	if e := s.SerializeString(&self.Name, MaxSizeName); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Health[0], MaxSizeStats); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Health[1], MaxSizeStats); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Mana[0], MaxSizeStats); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Mana[1], MaxSizeStats); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Stamina[0], MaxSizeStats); e != nil {
		return e
	}
	if e := s.SerializeInt(&self.Stamina[1], MaxSizeStats); e != nil {
		return e
	}
	return nil
}
