package list

import "testing"

func TestList_String(t *testing.T) {
	l := NewList[int]()
	for i := 1; i <= 2; i++ {
		l.PushBack(i)
	}
	if l.String() != "[1, 2]" {
		t.FailNow()
	}
}
