package treemap

import (
	"testing"
)

func TestTreeMap_Iter(t *testing.T) {
	tm := NewTreeMap[int, int]()
	for i := 0; i <= 2; i++ {
		tm.Set(i, i)
	}
	for iter := tm.Iterator(); iter != nil; iter.Next() {
		if iter.Key() != iter.Value() {
			t.FailNow()
		}
		if !iter.HasNext() {
			break
		}
	}
}
