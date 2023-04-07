package hashmap

import (
	"testing"
)

func TestHashMap_Iter(t *testing.T) {
	hm := NewHashMap[int, int]()
	for i := 0; i <= 2; i++ {
		hm.Set(i, i)
	}
	for iter := hm.Iterator(); iter != nil; iter.Next() {
		if iter.Key() != iter.Value() {
			t.FailNow()
		}
		if !iter.HasNext() {
			break
		}
	}
}
