package hashset

import "github.com/kkkunny/containers/hashmap"

type HashSetIter[T any] struct {
	iter *hashmap.HashMapIter[T, struct{}]
}

func (hs *HashSet[V]) Iterator() *HashSetIter[V] {
	if hs.Length() == 0 {
		return nil
	}
	return &HashSetIter[V]{iter: hs.data.Iterator()}
}

func (i *HashSetIter[V]) Value() V {
	return i.iter.Key()
}
func (i *HashSetIter[V]) HasNext() bool {
	return i.iter.HasNext()
}
func (i *HashSetIter[V]) Next() {
	i.iter.Next()
}

func HashSetMap[FV comparable, TV comparable](hs *HashSet[FV], fn func(v FV) TV, hasher func(TV) uint64) *HashSet[TV] {
	newHs := NewHashSetWithHasher(hasher)
	for iter := hs.Iterator(); iter != nil; iter.Next() {
		newHs.Add(fn(iter.Value()))
		if !iter.HasNext() {
			break
		}
	}
	return newHs
}
