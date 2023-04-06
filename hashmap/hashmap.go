package hashmap

import (
	"hash/fnv"
	"math/rand"
	"reflect"
	"unsafe"

	"github.com/kkkunny/containers/list"
)

const (
	defaultInitCapacity uint    = 16
	defaultLoadFactor   float64 = 0.75
)

type node[K, V any] struct {
	hash  uint64
	key   K
	value V
}

type HashMap[K, V any] struct {
	hasher func(K) uint64
	data   []*list.List[*node[K, V]]
	len    uint
}

func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return NewHashMapWithCapacity[K, V](defaultInitCapacity)
}
func NewHashMapWithCapacity[K comparable, V any](cap uint) *HashMap[K, V] {
	hasher := func(k K) uint64 {
		h := fnv.New64a()
		size := reflect.TypeOf(k).Size()
		ptr := uintptr(unsafe.Pointer(&k))
		for i := uintptr(0); i < size; i++ {
			b := *(*byte)(unsafe.Pointer(ptr + i))
			h.Write([]byte{b})
		}
		return h.Sum64()
	}
	return NewHashMapWith[K, V](cap, hasher)
}
func NewHashMapWithHasher[K, V any](hasher func(K) uint64) *HashMap[K, V] {
	return NewHashMapWith[K, V](defaultInitCapacity, hasher)
}
func NewHashMapWith[K, V any](cap uint, hasher func(K) uint64) *HashMap[K, V] {
	hm := &HashMap[K, V]{hasher: hasher}
	hm.initData(cap)
	return hm
}

func (hm *HashMap[K, V]) initData(cap uint) {
	hm.data = make([]*list.List[*node[K, V]], cap)
	for i := range hm.data {
		hm.data[i] = list.NewList[*node[K, V]]()
	}
}
func (hm *HashMap[K, V]) hash(k K) uint64 {
	return hm.hasher(k)
}
func (hm *HashMap[K, V]) index(hash uint64) uint {
	return uint(hash & uint64(len(hm.data)-1))
}
func (hm *HashMap[K, V]) expan() {
	oldData := hm.data
	hm.initData(uint(len(oldData)) * 2)
	for _, l := range oldData {
		for node := l.Front(); node != nil; node = node.Next() {
			value := node.Value()
			hm.insertNewData(value.key, value.value)
			hm.len--
		}
	}
}
func (hm *HashMap[K, V]) checkExpan() {
	if float64(hm.len)/float64(len(hm.data)) < defaultLoadFactor {
		return
	}
	hm.expan()
}
func (hm *HashMap[K, V]) Length() uint {
	return hm.len
}
func (hm *HashMap[K, V]) Capacity() uint {
	return uint(len(hm.data))
}
func (hm *HashMap[K, V]) find(hash uint64, k K) (*list.List[*node[K, V]], *list.ListNode[*node[K, V]]) {
	l := hm.data[hm.index(hash)]
	for node := l.Front(); node != nil; node = node.Next() {
		if hash == node.Value().hash {
			return l, node
		}
	}
	return l, nil
}
func (hm *HashMap[K, V]) insertNewData(k K, v V) bool {
	hash := hm.hash(k)
	l, _ := hm.find(hash, k)
	l.PushBack(&node[K, V]{
		hash:  hash,
		key:   k,
		value: v,
	})
	hm.len++
	return false
}
func (hm *HashMap[K, V]) Set(k K, v V) bool {
	hm.checkExpan()

	hash := hm.hash(k)

	l, n := hm.find(hash, k)
	if n != nil {
		n.Value().value = v
		return true
	}

	l.PushBack(&node[K, V]{
		hash:  hash,
		key:   k,
		value: v,
	})
	hm.len++
	return false
}
func (hm *HashMap[K, V]) Get(k K, v ...V) (V, bool) {
	hash := hm.hash(k)

	_, n := hm.find(hash, k)
	if n != nil {
		return n.Value().value, true
	}

	if len(v) > 0 {
		return v[len(v)-1], false
	}

	var val V
	return val, false
}
func (hm *HashMap[K, V]) ContainKey(k K) bool {
	hash := hm.hash(k)
	_, n := hm.find(hash, k)
	return n != nil
}
func (hm *HashMap[K, V]) Remove(k K, v ...V) (V, bool) {
	hash := hm.hash(k)

	l, n := hm.find(hash, k)
	if n != nil {
		hm.len--
		return l.RemoveNode(n).Value().value, true
	}

	if len(v) > 0 {
		return v[len(v)-1], false
	}

	var val V
	return val, false
}
func (hm *HashMap[K, V]) Clear() {
	hm.initData(defaultInitCapacity)
	hm.len = 0
}
func (hm *HashMap[K, V]) Keys() []K {
	if hm.len == 0 {
		return nil
	}

	keys := make([]K, hm.Length())
	var i int
	for _, l := range hm.data {
		for iter := l.Begin(); iter != nil; iter.Next() {
			keys[i] = iter.Value().key
			if !iter.HasNext() {
				break
			}
			i++
		}
	}
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	return keys
}
func (hm *HashMap[K, V]) Values() []V {
	if hm.len == 0 {
		return nil
	}

	vals := make([]V, hm.Length())
	var i int
	for _, l := range hm.data {
		for iter := l.Begin(); iter != nil; iter.Next() {
			vals[i] = iter.Value().value
			if !iter.HasNext() {
				break
			}
			i++
		}
	}
	rand.Shuffle(len(vals), func(i, j int) {
		vals[i], vals[j] = vals[j], vals[i]
	})

	return vals
}
