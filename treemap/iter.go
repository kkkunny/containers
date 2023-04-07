package treemap

import "github.com/kkkunny/containers/bstree"

type TreeMapIter[K, V any] struct {
	iter *bstree.BSTreeIter[entry[K, V]]
}

func (m *TreeMap[K, V]) Iterator() *TreeMapIter[K, V] {
	if m.tree.Length() == 0 {
		return nil
	}
	return &TreeMapIter[K, V]{iter: m.tree.Begin()}
}

func (i *TreeMapIter[K, V]) Key() K {
	return i.iter.Value().key
}
func (i *TreeMapIter[K, V]) Value() V {
	return i.iter.Value().value
}
func (i *TreeMapIter[K, V]) HasPrev() bool {
	return i.iter.HasPrev()
}
func (i *TreeMapIter[K, V]) Prev() {
	i.iter.Prev()
}
func (i *TreeMapIter[K, V]) HasNext() bool {
	return i.iter.HasNext()
}
func (i *TreeMapIter[K, V]) Next() {
	i.iter.Next()
}
