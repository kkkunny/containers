package heap

import "testing"

func TestHeap(t *testing.T) {
	h := NewMaxHeap[int]()
	h.Push(0)
	h.Push(1)
	if h.Length() != 2 {
		t.FailNow()
	}
	i := 1
	for h.Length() != 0 {
		if h.Pop() != i {
			t.FailNow()
		}
		i--
	}
}
