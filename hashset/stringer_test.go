package hashset

import "testing"

func TestHashSet_String(t *testing.T) {
	hs := NewHashSet[int]()
	for i := 1; i <= 2; i++ {
		hs.Add(i)
	}
	if  hs.String() != "{1, 2}" {
		t.FailNow()
	}
}
