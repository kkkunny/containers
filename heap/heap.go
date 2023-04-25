package heap

import (
	"github.com/kkkunny/containers/dynarray"
	"golang.org/x/exp/constraints"
)

type Heap[T any] struct {
	order func(T, T) int
	data  dynarray.DynArray[T]
	len   uint
}

func NewMinHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{
		order: func(l T, r T) int {
			if l < r {
				return -1
			} else if l == r {
				return 0
			} else {
				return 1
			}
		},
	}
}
func NewMaxHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{
		order: func(l T, r T) int {
			if l > r {
				return -1
			} else if l == r {
				return 0
			} else {
				return 1
			}
		},
	}
}
func NewHeapWith[T any](order func(T, T) int) *Heap[T] {
	return &Heap[T]{order: order}
}

func (h *Heap[T]) Length() uint {
	return h.len
}
func (h *Heap[T]) Push(v T) {
	h.data.Add(v)
	h.len++

	lastIndex := h.data.Length() - 1
	for i, fi := lastIndex, lastIndex/2; i != fi && h.order(h.data.Get(i), h.data.Get(fi)) < 0; i, fi = fi, fi/2 {
		iv, fiv := h.data.Get(i), h.data.Get(fi)
		h.data.Set(i, fiv)
		h.data.Set(fi, iv)
	}
}
func (h *Heap[T]) Peek() T {
	return h.data.Get(0)
}
func (h *Heap[T]) Pop() T {
	if h.len == 1 {
		h.len--
		return h.data.Remove(h.data.Length() - 1)
	}

	v := h.data.Set(0, h.data.Remove(h.data.Length()-1))
	h.len--

	lastIndex := h.data.Length() - 1
	for i, si := uint(0), uint(1); si <= lastIndex && h.order(h.data.Get(i), h.data.Get(si)) > 0; i, si = si, si*2+1 {
		iv, siv := h.data.Get(i), h.data.Get(si)
		h.data.Set(i, siv)
		h.data.Set(si, iv)
	}
	return v
}
func (h *Heap[T]) Clear() {
	h.data.Clear()
	h.len = 0
}
