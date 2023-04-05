package queue

import "github.com/kkkunny/containers/list"

type Queue[T any] struct {
	data *list.List[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: list.NewList[T]()}
}

func (q *Queue[T]) Length() uint {
	return q.data.Length()
}
func (q *Queue[T]) Push(v T) {
	q.data.PushBack(v)
}
func (q *Queue[T]) Peek() T {
	return q.data.Front().Value()
}
func (q *Queue[T]) Pop() T {
	return q.data.RemoveElem(q.data.Front()).Value()
}
func (q *Queue[T]) Clear() {
	q.data.Clear()
}
