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

func TestLinkedHashMap_Map(t *testing.T) {
	lhm := NewLinkedHashMap[int, int]()
	for i := 0; i <= 2; i++ {
		lhm.Set(i, i)
	}
	newLhm := LinkedHashMapMap(lhm, func(k int, v int) (int, int) {
		return k + 1, v + 1
	})
	for iter := newLhm.Begin(); iter != nil; iter.Next() {
		if iter.Key() != iter.Value() {
			t.FailNow()
		}
		if !iter.HasNext() {
			break
		}
	}
}
