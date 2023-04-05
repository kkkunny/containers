package stack

import "github.com/kkkunny/containers/list"

type StackIter[T any] struct {
	iter *list.ListIter[T]
}

func (s *Stack[T]) Iterator() *StackIter[T] {
	if s.Length() == 0 {
		return nil
	}
	return &StackIter[T]{iter: s.data.End()}
}

func (i *StackIter[T]) Value() T {
	return i.iter.Value()
}
func (i *StackIter[T]) HasNext() bool {
	return i.iter.HasPrev()
}
func (i *StackIter[T]) Next() {
	i.iter.Prev()
}

func StackMap[From, To any](s *Stack[From], fn func(v From) To) *Stack[To] {
	ns := NewStack[To]()
	for iter := s.data.Begin(); iter != nil; iter.Next() {
		ns.data.PushBack(fn(iter.Value()))
		if !iter.HasNext() {
			break
		}
	}
	return ns
}
