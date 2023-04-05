package queue

import "testing"

func TestQueue_Iter(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i <= 2; i++ {
		q.Push(i)
	}
	var i int
	for iter := q.Iterator(); iter != nil; iter.Next() {
		if i != iter.Value() {
			t.FailNow()
		}
		i++
		if !iter.HasNext() {
			break
		}
	}
}

func TestQueue_Map(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i <= 2; i++ {
		q.Push(i)
	}
	nq := QueueMap(q, func(v int) int {
		return v + 1
	})
	var i int
	for iter := nq.Iterator(); iter != nil; iter.Next() {
		if i != iter.Value()-1 {
			t.FailNow()
		}
		i++
		if !iter.HasNext() {
			break
		}
	}
}
