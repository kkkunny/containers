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
