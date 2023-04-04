package dynarray

import "testing"

func TestDynArray_String(t *testing.T) {
	da := NewDynArray[int]()
	for i := 1; i <= 2; i++ {
		da.Add(i)
	}
	if da.String() != "[1, 2]" {
		t.FailNow()
	}
}
