package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	q.Push(0)
	q.Push(1)
	if q.Length() != 2 {
		t.FailNow()
	}
	var i int
	for q.Length() != 0 {
		if q.Pop() != i {
			t.FailNow()
		}
		i++
	}
}
