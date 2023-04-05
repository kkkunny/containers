package hashset

import "github.com/kkkunny/containers/hashmap"

type HashSet[T comparable] struct {
	data *hashmap.HashMap[T, struct{}]
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{data: hashmap.NewHashMap[T, struct{}]()}
}
func NewHashSetWith[T comparable](cap uint) *HashSet[T] {
	return &HashSet[T]{data: hashmap.NewHashMapWith[T, struct{}](cap)}
}

func (hs *HashSet[T]) Length() uint {
	return hs.data.Length()
}
func (hs *HashSet[T]) Add(v T) bool {
	if hs.data.ContainKey(v) {
		return false
	}
	hs.data.Set(v, struct{}{})
	return true
}
func (hs *HashSet[T]) Remove(v T) bool {
	if !hs.data.ContainKey(v) {
		return false
	}
	hs.data.Remove(v)
	return true
}
func (hs *HashSet[T]) Clear() {
	hs.data.Clear()
}
