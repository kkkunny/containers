package stack

import "testing"

func TestStack_String(t *testing.T) {
	s := NewStack[int]()
	for i := 1; i <= 2; i++ {
		s.Push(i)
	}
	if s.String() != "[1, 2]" {
		t.FailNow()
	}
}
