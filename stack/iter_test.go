package stack

import "testing"

func TestStack_Iter(t *testing.T) {
	s := NewStack[int]()
	for i := 0; i <= 2; i++ {
		s.Push(i)
	}
	i := 2
	for iter := s.Iterator(); iter != nil; iter.Next() {
		if i != iter.Value() {
			t.FailNow()
		}
		i--
		if !iter.HasNext() {
			break
		}
	}
}

func TestStack_Map(t *testing.T) {
	s := NewStack[int]()
	for i := 0; i <= 2; i++ {
		s.Push(i)
	}
	ns := StackMap(s, func(v int) int {
		return v + 1
	})
	i := 2
	for iter := ns.Iterator(); iter != nil; iter.Next() {
		if i != iter.Value()-1 {
			t.FailNow()
		}
		i--
		if !iter.HasNext() {
			break
		}
	}
}
