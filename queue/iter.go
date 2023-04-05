package queue

import "github.com/kkkunny/containers/list"

type QueueIter[T any] struct {
	iter *list.ListIter[T]
}

func (q *Queue[T]) Iterator() *QueueIter[T] {
	if q.Length() == 0 {
		return nil
	}
	return &QueueIter[T]{iter: q.data.Begin()}
}

func (i *QueueIter[T]) Value() T {
	return i.iter.Value()
}
func (i *QueueIter[T]) HasNext() bool {
	return i.iter.HasNext()
}
func (i *QueueIter[T]) Next() {
	i.iter.Next()
}

func QueueMap[From, To any](q *Queue[From], fn func(v From) To) *Queue[To] {
	nq := NewQueue[To]()
	for iter := q.Iterator(); iter != nil; iter.Next() {
		nq.Push(fn(iter.Value()))
		if !iter.HasNext() {
			break
		}
	}
	return nq
}
