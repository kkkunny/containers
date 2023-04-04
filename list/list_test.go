package list

import "testing"

func TestNewList(t *testing.T) {
	l := NewList[int]()
	if l.Length() != 0 {
		t.FailNow()
	}
}

func TestList_Push(t *testing.T) {
	l := NewList[int]()
	l.PushFront(1)
	l.PushBack(2)
	if l.Length() != 2 {
		t.FailNow()
	}
	var i int
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		if i == 0 {
			if iter.Value() != 1 {
				t.FailNow()
			}
		} else if i == 1 {
			if iter.Value() != 2 {
				t.FailNow()
			}
		}
		i++
	}
}

func TestList_Move(t *testing.T) {
	l := NewList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.MoveToBack(l.Front())
	var i int
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		if i == 0 {
			if iter.Value() != 2 {
				t.FailNow()
			}
		} else if i == 1 {
			if iter.Value() != 1 {
				t.FailNow()
			}
		}
		i++
	}
}
