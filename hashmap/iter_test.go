package hashmap

import "testing"

func TestHashMap_Iter(t *testing.T) {
	hm := NewHashMap[int, int]()
	for i := 0; i <= 2; i++ {
		hm.Set(i, i)
	}
	for iter := hm.Iterator(); iter != nil; iter.Next() {
		if iter.Key() != iter.Value() {
			t.FailNow()
		}
		if !iter.HasNext(){
			break
		}
	}
}

func TestHashMap_Map(t *testing.T) {
	hm := NewHashMap[int, int]()
	for i := 0; i <= 2; i++ {
		hm.Set(i, i)
	}
	newHm := HashMapMap(hm, func(k int, v int) (int, int) {
		return k+1, v+1
	})
	for iter := newHm.Iterator(); iter != nil; iter.Next() {
		if iter.Key() != iter.Value() {
			t.FailNow()
		}
		if !iter.HasNext(){
			break
		}
	}
}
