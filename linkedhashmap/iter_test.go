package linkedhashmap

import "testing"

func TestLinkedHashMap_Iter(t *testing.T) {
	lhm := NewLinkedHashMap[int, int]()
	for i := 0; i <= 2; i++ {
		lhm.Set(i, i)
	}
	for iter := lhm.Begin(); iter != nil; iter.Next() {
		if iter.Key() != iter.Value() {
			t.FailNow()
		}
		if !iter.HasNext() {
			break
		}
	}
}
