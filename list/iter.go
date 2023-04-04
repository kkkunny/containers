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
func (i *ListIter[T]) Prev() {
	if !i.HasPrev() {
		panic("index out of range")
	}
	i.elem = i.elem.Prev()
}
func (i *ListIter[T]) HasNext() bool {
	return i.elem.Next() != nil
}
func (i *ListIter[T]) Next() {
	if !i.HasNext() {
		panic("index out of range")
	}
	i.elem = i.elem.Next()
}

func ListMap[From, To any](l *List[From], fn func(v From)To)*List[To]{
	nl := NewList[To]()
	for iter:=l.Begin(); iter != nil; iter.Next() {
		nl.PushBack(fn(iter.Value()))
		if !iter.HasNext(){
			break
		}
	}
	return nl
}
