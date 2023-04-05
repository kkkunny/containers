package list

type Elem[T any] struct {
	prev, next *Elem[T]
	val        T
	list       *List[T]
}

func (e *Elem[T]) Prev() *Elem[T] {
	if e.prev == e.list.root {
		return nil
	}
	return e.prev
}
func (e *Elem[T]) Next() *Elem[T] {
	if e.next == e.list.root {
		return nil
	}
	return e.next
}
func (e *Elem[T]) Value() T {
	return e.val
}

type List[T any] struct {
	root *Elem[T]
	len  uint
}

func NewList[T any]() *List[T] {
	l := new(List[T])
	l.root = &Elem[T]{list: l}
	l.root.prev = l.root
	l.root.next = l.root
	return l
}

func (l *List[T]) Length() uint {
	return l.len
}
func (l *List[T]) checkElem(e *Elem[T]) {
	if e.list != l {
		panic("the element is not in this list")
	}
}
func (l *List[T]) RemoveElem(e *Elem[T]) T {
	l.checkElem(e)
	if e.prev != nil {
		e.prev.next = e.next
	}
	if e.next != nil {
		e.next.prev = e.prev
	}
	l.len--
	return e.val
}
func (l *List[T]) MoveToFrontOfElem(elem, target *Elem[T]) T {
	l.checkElem(elem)
	l.checkElem(target)

	l.RemoveElem(elem)
	l.len++

	target.prev.next = elem
	elem.prev = target.prev

	target.prev = elem
	elem.next = target
	return elem.val
}
func (l *List[T]) MoveToBackOfElem(elem, target *Elem[T]) T {
	l.checkElem(elem)
	l.checkElem(target)

	l.RemoveElem(elem)
	l.len++

	target.next.prev = elem
	elem.next = target.next

	target.next = elem
	elem.prev = target
	return elem.val
}
func (l *List[T]) MoveToFront(e *Elem[T]) T {
	return l.MoveToBackOfElem(e, l.root)
}
func (l *List[T]) MoveToBack(e *Elem[T]) T {
	return l.MoveToFrontOfElem(e, l.root)
}
func (l *List[T]) PushFront(v T) T {
	l.len++
	elem := &Elem[T]{val: v, list: l}
	return l.MoveToFront(elem)
}
func (l *List[T]) PushBack(v T) T {
	l.len++
	elem := &Elem[T]{val: v, list: l}
	return l.MoveToBack(elem)
}
func (l *List[T]) Front() *Elem[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}
func (l *List[T]) Back() *Elem[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}
func (l *List[T]) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.len = 0
}
