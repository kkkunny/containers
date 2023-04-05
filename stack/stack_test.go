package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(0)
	if s.Length() != 2 {
		t.FailNow()
	}
	var i int
	for s.Length() != 0 {
		if s.Pop() != i {
			t.FailNow()
		}
		i++
	}
}
