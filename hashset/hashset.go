package hashset

import "github.com/kkkunny/containers/hashmap"

type HashSet[T any] struct {
	data *hashmap.HashMap[T, struct{}]
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{data: hashmap.NewHashMap[T, struct{}]()}
}
func NewHashSetWithCapacity[T comparable](cap uint) *HashSet[T] {
	return &HashSet[T]{data: hashmap.NewHashMapWithCapacity[T, struct{}](cap)}
}
func NewHashSetWithHasher[T any](hasher func(T) uint64) *HashSet[T] {
	return &HashSet[T]{data: hashmap.NewHashMapWithHasher[T, struct{}](hasher)}
}
func NewHashSetWith[T any](cap uint, hasher func(T) uint64) *HashSet[T] {
	return &HashSet[T]{data: hashmap.NewHashMapWith[T, struct{}](cap, hasher)}
}

func (hs *HashSet[T]) Length() uint {
	return hs.data.Length()
}
func (hs *HashSet[T]) Add(v T) bool {
	return !hs.data.Set(v, struct{}{})
}
func (hs *HashSet[T]) Remove(v T) bool {
	_, ok := hs.data.Remove(v)
	return ok
}
func (hs *HashSet[T]) Clear() {
	hs.data.Clear()
}
