package hashmap

import "testing"

func TestHashMap_String(t *testing.T) {
	hm := NewHashMap[int, int]()
	for i := 1; i <= 2; i++ {
		hm.Set(i, i)
	}
	if hm.String() != "{1: 1, 2: 2}" {
		t.FailNow()
	}
}
