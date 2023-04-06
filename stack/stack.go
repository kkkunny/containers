package stack

import "github.com/kkkunny/containers/list"

type Stack[T any] struct {
	data *list.List[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: list.NewList[T]()}
}

func (s *Stack[T]) Length() uint {
	return s.data.Length()
}
func (s *Stack[T]) Push(v T) {
	s.data.PushBack(v)
}
func (s *Stack[T]) Peek() T {
	return s.data.Back().Value()
}
func (s *Stack[T]) Pop() T {
	return s.data.RemoveNode(s.data.Back()).Value()
}
func (s *Stack[T]) Clear() {
	s.data.Clear()
}
