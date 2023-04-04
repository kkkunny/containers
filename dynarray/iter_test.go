package dynarray

import "testing"

func TestDynArray_Iter(t *testing.T) {
	da := NewDynArray[int]()
	for i := 0; i <= 2; i++ {
		da.Add(i)
	}
	for iter := da.Begin(); iter != nil; iter.Next() {
		if int(iter.Index()) != iter.Value() {
			t.FailNow()
		}
		if !iter.HasNext(){
			break
		}
	}
}

func TestDynArray_Map(t *testing.T) {
	da := NewDynArray[int]()
	for i := 0; i <= 2; i++ {
		da.Add(i)
	}
	newDa := DynArrayMap(da, func(i uint, v int) int {
		return v + 1
	})
	for iter := newDa.Begin(); iter != nil; iter.Next() {
		if int(iter.Index()) != iter.Value() - 1 {
			t.FailNow()
		}
		if !iter.HasNext(){
			break
		}
	}
}
