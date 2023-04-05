package linkedhashmap

import "testing"

func TestLinkedHashMap_String(t *testing.T) {
	lhm := NewLinkedHashMap[int, int]()
	for i := 1; i <= 2; i++ {
		lhm.Set(i, i)
	}
	if lhm.String() != "{1: 1, 2: 2}" {
		t.FailNow()
	}
}
