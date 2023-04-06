package hashset

import "testing"

func TestHashSet(t *testing.T) {
	hs := NewHashSetWithCapacity[int](10)
	for i := 0; i < 10; i++ {
		hs.Add(i)
	}
	if hs.Length() != 10 {
		t.FailNow()
	}
	if hs.Add(1) {
		t.FailNow()
	}
	if !hs.Add(11) {
		t.FailNow()
	}
}
