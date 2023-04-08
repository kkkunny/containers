package priorityqueue

import "testing"

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int, int]()
	q.Push(0, 0)
	q.Push(1, 1)
	if q.Length() != 2 {
		t.FailNow()
	}
	i := 1
	for q.Length() != 0 {
		if _, v := q.Pop(); v != i {
			t.FailNow()
		}
		i--
	}
}
