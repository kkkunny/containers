package heap

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewMaxHeap[int]()
	h.Push(0)
	h.Push(1)
	if h.Length() != 2 {
		t.FailNow()
	}
	i := 1
	for h.Length() != 0 {
		if h.Pop() != i {
			t.FailNow()
		}
		i--
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkPush_Heap(b *testing.B) {
	h := NewMinHeap[int]()
	for i := 10000; i > 0; i-- {
		h.Push(i)
	}
}

func BenchmarkPush_Stdlib(b *testing.B) {
	var h IntHeap
	heap.Init(&h)
	for i := 1; i <= 10000; i++ {
		h.Push(i)
	}
}
