package linkedhashmap

import "github.com/kkkunny/containers/list"

type LinkedHashMapIter[K any, V any] struct {
	iter *list.ListIter[linkedHashMapElem[K, V]]
}

func (lhm *LinkedHashMap[K, V]) Begin() *LinkedHashMapIter[K, V] {
	if lhm.Length() == 0 {
		return nil
	}
	return &LinkedHashMapIter[K, V]{iter: lhm.data.Begin()}
}
func (lhm *LinkedHashMap[K, V]) End() *LinkedHashMapIter[K, V] {
	if lhm.Length() == 0 {
		return nil
	}
	return &LinkedHashMapIter[K, V]{iter: lhm.data.End()}
}

func (i *LinkedHashMapIter[K, V]) Key() K {
	return i.iter.Value().key
}
func (i *LinkedHashMapIter[K, V]) Value() V {
	return i.iter.Value().value
}
func (i *LinkedHashMapIter[K, V]) HasPrev() bool {
	return i.iter.HasPrev()
}
func (i *LinkedHashMapIter[K, V]) Prev() {
	i.iter.Prev()
}
func (i *LinkedHashMapIter[K, V]) HasNext() bool {
	return i.iter.HasNext()
}
func (i *LinkedHashMapIter[K, V]) Next() {
	i.iter.Next()
}
