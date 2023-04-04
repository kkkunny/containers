package list

type ListIter[T any] struct {
	elem *Elem[T]
}

func (l *List[T]) Begin() *ListIter[T] {
	if l.Length() == 0 {
		return nil
	}
	return &ListIter[T]{elem: l.Front()}
}
func (l *List[T]) End() *ListIter[T] {
	if l.Length() == 0 {
		return nil
	}
	return &ListIter[T]{elem: l.Back()}
}

func (i *ListIter[T]) Value() T {
	return i.elem.Value()
}
func (i *ListIter[T]) HasPrev() bool {
	return i.elem.Prev() != nil
}
func (i *ListIter[T]) Prev() *ListIter[T] {
	if !i.HasPrev() {
		return nil
	}
	return &ListIter[T]{elem: i.elem.Prev()}
}
func (i *ListIter[T]) HasNext() bool {
	return i.elem.Next() != nil
}
func (i *ListIter[T]) Next() *ListIter[T] {
	if !i.HasNext() {
		return nil
	}
	return &ListIter[T]{elem: i.elem.Next()}
}
