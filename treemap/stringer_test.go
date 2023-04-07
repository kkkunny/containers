package treemap

import "testing"

func TestTreeMap_String(t *testing.T) {
	tm := NewTreeMap[int, int]()
	tm.Set(5, 5)
	tm.Set(1, 1)
	tm.Set(7, 7)
	tm.Set(3, 3)
	tm.Set(9, 9)
	tm.Set(0, 0)
	tm.Set(2, 2)
	tm.Set(4, 4)
	tm.Set(6, 6)
	tm.Set(8, 8)
	if tm.String() != "{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}" {
		t.FailNow()
	}
}
