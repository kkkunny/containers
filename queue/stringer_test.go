package queue

import "testing"

func TestQueue_String(t *testing.T) {
	q := NewQueue[int]()
	for i := 1; i <= 2; i++ {
		q.Push(i)
	}
	if q.String() != "[1, 2]" {
		t.FailNow()
	}
}
