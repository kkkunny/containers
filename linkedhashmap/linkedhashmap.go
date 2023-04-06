package linkedhashmap

import (
	"github.com/kkkunny/containers/hashmap"
	"github.com/kkkunny/containers/list"
)

type linkedHashMapElem[K comparable, V any] struct {
	key   K
	value V
}

type LinkedHashMap[K comparable, V any] struct {
	hash *hashmap.HashMap[K, *list.ListNode[linkedHashMapElem[K, V]]]
	data *list.List[linkedHashMapElem[K, V]]
}

func NewLinkedHashMap[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hash: hashmap.NewHashMap[K, *list.ListNode[linkedHashMapElem[K, V]]](),
		data: list.NewList[linkedHashMapElem[K, V]](),
	}
}
func NewLinkedHashMapWith[K comparable, V any](cap uint) *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hash: hashmap.NewHashMapWith[K, *list.ListNode[linkedHashMapElem[K, V]]](cap),
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
	listElem := lhm.hash.Get(k)
	if listElem == nil && len(v) > 0 {
		return v[len(v)-1]
	} else if listElem == nil {
		var val V
		return val
	}
	return listElem.Value().value
}
func (lhm *LinkedHashMap[K, V]) ContainKey(k K) bool {
	return lhm.hash.ContainKey(k)
}
func (lhm *LinkedHashMap[K, V]) Remove(k K, v ...V) V {
	listElem := lhm.hash.Remove(k)
	if listElem == nil && len(v) > 0 {
		return v[len(v)-1]
	} else if listElem == nil {
		var val V
		return val
	}
	return lhm.data.RemoveElem(listElem).Value().value
}
func (lhm *LinkedHashMap[K, V]) Clear() {
	lhm.hash.Clear()
	lhm.data.Clear()
}
func (lhm *LinkedHashMap[K, V]) Front() (K, V) {
	frontListElem := lhm.data.Front()
	if frontListElem == nil {
		var key K
		var val V
		return key, val
	}
	frontElem := frontListElem.Value()
	return frontElem.key, frontElem.value
}
func (lhm *LinkedHashMap[K, V]) Back() (K, V) {
	backListElem := lhm.data.Back()
	if backListElem == nil {
		var key K
		var val V
		return key, val
	}
	backElem := backListElem.Value()
	return backElem.key, backElem.value
}
