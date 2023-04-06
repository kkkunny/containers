package list

type ListNode[T any] struct {
	prev, next *ListNode[T]
	list       *List[T]
	val        T
}

func (n *ListNode[T]) Prev() *ListNode[T] {
	if n.prev == n.list.root {
		return nil
	}
	return n.prev
}
func (n *ListNode[T]) Next() *ListNode[T] {
	if n.next == n.list.root {
		return nil
	}
	return n.next
}
func (n *ListNode[T]) Value() T {
	return n.val
}

type List[T any] struct {
	root *ListNode[T]
	len  uint
}

func NewList[T any]() *List[T] {
	l := new(List[T])
	l.root = &ListNode[T]{list: l}
	l.root.prev = l.root
	l.root.next = l.root
	return l
}

func (l *List[T]) Length() uint {
	return l.len
}
func (l *List[T]) checkNode(e *ListNode[T]) {
	if e.list != l {
		panic("the element is not in this list")
	}
}
func (l *List[T]) RemoveNode(e *ListNode[T]) *ListNode[T] {
	l.checkNode(e)
	e.list = nil
	if e.prev != nil {
		e.prev.next = e.next
	}
	if e.next != nil {
		e.next.prev = e.prev
	}
	l.len--
	return e
}
func (l *List[T]) MoveToFrontOfNode(elem, target *ListNode[T]) *ListNode[T] {
	l.checkNode(elem)
	l.checkNode(target)

	l.RemoveNode(elem)
	elem.list = l
	l.len++

	target.prev.next = elem
	elem.prev = target.prev

	target.prev = elem
	elem.next = target
	return elem
}
func (l *List[T]) MoveToBackOfNode(elem, target *ListNode[T]) *ListNode[T] {
	l.checkNode(elem)
	l.checkNode(target)

	l.RemoveNode(elem)
	elem.list = l
	l.len++

	target.next.prev = elem
	elem.next = target.next

	target.next = elem
	elem.prev = target
	return elem
}
func (l *List[T]) MoveToFront(e *ListNode[T]) *ListNode[T] {
	return l.MoveToBackOfNode(e, l.root)
}
func (l *List[T]) MoveToBack(e *ListNode[T]) *ListNode[T] {
	return l.MoveToFrontOfNode(e, l.root)
}
func (l *List[T]) PushFront(v T) *ListNode[T] {
	l.len++
	elem := &ListNode[T]{val: v, list: l}
	return l.MoveToFront(elem)
}
func (l *List[T]) PushBack(v T) *ListNode[T] {
	l.len++
	elem := &ListNode[T]{val: v, list: l}
	return l.MoveToBack(elem)
}
func (l *List[T]) Front() *ListNode[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}
func (l *List[T]) Back() *ListNode[T] {
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
