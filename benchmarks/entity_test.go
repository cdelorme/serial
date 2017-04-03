package benchmarks

import (
	"bytes"
	"testing"

	"github.com/cdelorme/serial"
)

func TestEntity(t *testing.T) {
	var b bytes.Buffer
	r, w := serial.NewReader(&b), serial.NewWriter(&b)
	o, i := *NewEntity(), Entity{}

	if e := i.Serialize(r); e == nil {
		t.Error("failed to capture error with no data...")
	}
	if e := o.Serialize(w); e != nil || b.Len() == 0 {
		if e != nil {
			t.Logf("Error: %s\n", e)
		}
		t.Errorf("failed to write structure (%T)...\n", o)
	}
	if e := i.Serialize(r); e != nil || o.Name != i.Name || o.Health[0] != i.Health[0] || o.Health[1] != i.Health[1] || len(i.Friends) != len(o.Friends) || len(i.Statuses) != len(o.Statuses) {
		t.Errorf("\nfailed to read structure %T...\n\n", o)
		if e != nil {
			t.Logf("\nError: %s\n\n", e)
		}
	}
}
