package linkedhashmap

import (
	"github.com/kkkunny/containers/hashmap"
	"github.com/kkkunny/containers/list"
)

type elem[K any, V any] struct {
	key   K
	value V
}

type LinkedHashMap[K any, V any] struct {
	hash *hashmap.HashMap[K, *list.ListNode[elem[K, V]]]
	data *list.List[elem[K, V]]
}

func NewLinkedHashMap[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hash: hashmap.NewHashMap[K, *list.ListNode[elem[K, V]]](),
		data: list.NewList[elem[K, V]](),
	}
}
func NewLinkedHashMapWithCapacity[K comparable, V any](cap uint) *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hash: hashmap.NewHashMapWithCapacity[K, *list.ListNode[elem[K, V]]](cap),
		data: list.NewList[elem[K, V]](),
	}
}
func NewLinkedHashMapWithHasher[K any, V any](hasher func(K) uint64) *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hash: hashmap.NewHashMapWithHasher[K, *list.ListNode[elem[K, V]]](hasher),
		data: list.NewList[elem[K, V]](),
	}
}
func NewLinkedHashMapWith[K any, V any](cap uint, hasher func(K) uint64) *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hash: hashmap.NewHashMapWith[K, *list.ListNode[elem[K, V]]](cap, hasher),
		data: list.NewList[elem[K, V]](),
	}
}

func (lhm *LinkedHashMap[K, V]) Length() uint {
	return lhm.data.Length()
}
func (lhm *LinkedHashMap[K, V]) Capacity() uint {
	return lhm.hash.Capacity()
}
func (lhm *LinkedHashMap[K, V]) Set(k K, v V) bool {
	_, ok := lhm.Remove(k)
	elem := elem[K, V]{
		key:   k,
		value: v,
	}
	lhm.hash.Set(k, lhm.data.PushBack(elem))
	return ok
}
func (lhm *LinkedHashMap[K, V]) Get(k K, v ...V) (V, bool) {
	node, _ := lhm.hash.Get(k)
	if node == nil && len(v) > 0 {
		return v[len(v)-1], false
	} else if node == nil {
		var val V
		return val, false
	}
	return node.Value().value, true
}
func (lhm *LinkedHashMap[K, V]) ContainKey(k K) bool {
	return lhm.hash.ContainKey(k)
}
func (lhm *LinkedHashMap[K, V]) Remove(k K, v ...V) (V, bool) {
	node, _ := lhm.hash.Remove(k)
	if node == nil && len(v) > 0 {
		return v[len(v)-1], false
	} else if node == nil {
		var val V
		return val, false
	}
	return lhm.data.RemoveNode(node).Value().value, true
}
func (lhm *LinkedHashMap[K, V]) Clear() {
	lhm.hash.Clear()
	lhm.data.Clear()
}
func (lhm *LinkedHashMap[K, V]) Front() (K, V) {
	frontNode := lhm.data.Front()
	if frontNode == nil {
		var key K
		var val V
		return key, val
	}
	frontElem := frontNode.Value()
	return frontElem.key, frontElem.value
}
func (lhm *LinkedHashMap[K, V]) Back() (K, V) {
	backNode := lhm.data.Back()
	if backNode == nil {
		var key K
		var val V
		return key, val
	}
	backElem := backNode.Value()
	return backElem.key, backElem.value
}
func (lhm *LinkedHashMap[K, V]) Keys() []K {
	if lhm.data.Length() == 0 {
		return nil
	}

	keys := make([]K, lhm.data.Length())
	var i int
	for iter := lhm.data.Begin(); iter != nil; iter.Next() {
		keys[i] = iter.Value().key
		if !iter.HasNext() {
			break
		}
		i++
	}

	return keys
}
func (lhm *LinkedHashMap[K, V]) Values() []V {
	if lhm.data.Length() == 0 {
		return nil
	}

	vals := make([]V, lhm.data.Length())
	var i int
	for iter := lhm.data.Begin(); iter != nil; iter.Next() {
		vals[i] = iter.Value().value
		if !iter.HasNext() {
			break
		}
		i++
	}

	return vals
}
