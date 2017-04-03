package benchmarks

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func NewEntity() *Entity {
	e := &Entity{
		Name:    randString((rand.Intn(12) + 8)),
		Health:  [2]uint16{uint16(rand.Intn(100)), uint16(rand.Intn(100))},
		Mana:    [2]uint16{uint16(rand.Intn(100)), uint16(rand.Intn(100))},
		Stamina: [2]uint16{uint16(rand.Intn(100)), uint16(rand.Intn(100))},
		Dead:    rand.Intn(1) != 0,
	}

	f := rand.Intn(20)
	for i := 0; i < f; i++ {
		e.Friends = append(e.Friends, rand.Int63())
	}

	s := rand.Intn(5)
	for i := 0; i < s; i++ {
		st := NewStatus()
		e.Statuses = append(e.Statuses, *st)
	}

	return e
}

func NewStatus() *Status {
	return &Status{
		Name:     randString((rand.Intn(12) + 8)),
		Duration: uint16(rand.Intn(300)),
	}
}
