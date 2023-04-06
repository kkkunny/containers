package linkedhashmap

import (
	"github.com/kkkunny/containers/hashmap"
	"github.com/kkkunny/containers/list"
)

type linkedHashMapElem[K any, V any] struct {
	key   K
	value V
}

type LinkedHashMap[K any, V any] struct {
	hash *hashmap.HashMap[K, *list.ListNode[linkedHashMapElem[K, V]]]
	data *list.List[linkedHashMapElem[K, V]]
}

func NewLinkedHashMap[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hash: hashmap.NewHashMap[K, *list.ListNode[linkedHashMapElem[K, V]]](),
		data: list.NewList[linkedHashMapElem[K, V]](),
	}
}
func NewLinkedHashMapWithCapacity[K comparable, V any](cap uint) *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hash: hashmap.NewHashMapWithCapacity[K, *list.ListNode[linkedHashMapElem[K, V]]](cap),
		data: list.NewList[linkedHashMapElem[K, V]](),
	}
}
func NewLinkedHashMapWithHasher[K any, V any](hasher func(K) uint64) *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hash: hashmap.NewHashMapWithHasher[K, *list.ListNode[linkedHashMapElem[K, V]]](hasher),
		data: list.NewList[linkedHashMapElem[K, V]](),
	}
}
func NewLinkedHashMapWith[K any, V any](cap uint, hasher func(K) uint64) *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hash: hashmap.NewHashMapWith[K, *list.ListNode[linkedHashMapElem[K, V]]](cap, hasher),
		data: list.NewList[linkedHashMapElem[K, V]](),
	}
}

func (lhm *LinkedHashMap[K, V]) Length() uint {
	return lhm.data.Length()
}
func (lhm *LinkedHashMap[K, V]) Set(k K, v V) bool {
	if lhm.hash.ContainKey(k) {
		return false
	}
	elem := linkedHashMapElem[K, V]{
		key:   k,
		value: v,
	}
	lhm.hash.Set(k, lhm.data.PushBack(elem))
	return true
}
func (lhm *LinkedHashMap[K, V]) Get(k K, v ...V) V {
	node, _ := lhm.hash.Get(k)
	if node == nil && len(v) > 0 {
		return v[len(v)-1]
	} else if node == nil {
		var val V
		return val
	}
	return node.Value().value
}
func (lhm *LinkedHashMap[K, V]) ContainKey(k K) bool {
	return lhm.hash.ContainKey(k)
}
func (lhm *LinkedHashMap[K, V]) Remove(k K, v ...V) V {
	node, _ := lhm.hash.Remove(k)
	if node == nil && len(v) > 0 {
		return v[len(v)-1]
	} else if node == nil {
		var val V
		return val
	}
	return lhm.data.RemoveNode(node).Value().value
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
