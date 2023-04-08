package priorityqueue

import (
	"github.com/kkkunny/containers/heap"
	"github.com/kkkunny/containers/list"
	"golang.org/x/exp/constraints"
)

type elem[P, V any] struct {
	priority P
	value    V
}

type PriorityQueue[P, V any] struct {
	sorter *heap.Heap[*list.ListNode[elem[P, V]]]
	data   *list.List[elem[P, V]]
}

func NewPriorityQueue[P constraints.Ordered, V any]() *PriorityQueue[P, V] {
	return &PriorityQueue[P, V]{
		sorter: heap.NewHeapWith[*list.ListNode[elem[P, V]]](
			func(l *list.ListNode[elem[P, V]], r *list.ListNode[elem[P, V]]) int {
				lv, rv := l.Value().priority, r.Value().priority
				if lv > rv {
					return -1
				} else if lv == rv {
					return 0
				} else {
					return 1
				}
			},
		),
		data: list.NewList[elem[P, V]](),
	}
}
func NewPriorityQueueWith[P, V any](order func(P, P) int) *PriorityQueue[P, V] {
	return &PriorityQueue[P, V]{
		sorter: heap.NewHeapWith[*list.ListNode[elem[P, V]]](
			func(l *list.ListNode[elem[P, V]], r *list.ListNode[elem[P, V]]) int {
				return 0 - order(l.Value().priority, r.Value().priority)
			},
		),
		data: list.NewList[elem[P, V]](),
	}
}

func (q *PriorityQueue[P, V]) Length() uint {
	return q.data.Length()
}
func (q *PriorityQueue[P, V]) Push(p P, v V) {
	node := q.data.PushBack(
		elem[P, V]{
			priority: p,
			value:    v,
		},
	)
	q.sorter.Push(node)
}
func (q *PriorityQueue[P, V]) Peek() (P, V) {
	elem := q.sorter.Peek().Value()
	return elem.priority, elem.value
}
func (q *PriorityQueue[P, V]) Pop() (P, V) {
	elem := q.data.RemoveNode(q.sorter.Pop()).Value()
	return elem.priority, elem.value
}
func (q *PriorityQueue[P, V]) Clear() {
	q.sorter.Clear()
	q.data.Clear()
}
