package heap

import "golang.org/x/exp/constraints"

type Heap[T any] struct {
	order func(T, T) int
	data  []T
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

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *Heap[T]) less(i, j int) bool {
	return h.order(h.data[i], h.data[j]) < 0
}
func (h *Heap[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}
func (h *Heap[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}
func (h *Heap[T]) Length() uint {
	return h.len
}
func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.up(len(h.data) - 1)
	h.len++
}
func (h *Heap[T]) Peek() T {
	return h.data[len(h.data)-1]
}
func (h *Heap[T]) Pop() T {
	n := len(h.data) - 1
	h.swap(0, n)
	h.down(0, n)

	v := h.data[n]
	h.data = h.data[:n]
	h.len--
	return v
}
func (h *Heap[T]) Clear() {
	h.data = make([]T, 0)
	h.len = 0
}
