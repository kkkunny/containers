package hashmap

import (
	"math/rand"
)

const (
	defaultInitCapacity uint    = 16
	defaultLoadFactor   float64 = 0.75
)

type entry[K, V any] struct {
	next  *entry[K, V]
	hash  uint64
	key   K
	value V
}

func newEntry[K, V any](hash uint64, k K, v V) *entry[K, V] {
	return &entry[K, V]{
		hash:  hash,
		key:   k,
		value: v,
	}
}

type HashMap[K, V any] struct {
	hasher func(K) uint64
	data   []entry[K, V]
	len    uint
}

func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return NewHashMapWithCapacity[K, V](defaultInitCapacity)
}
func NewHashMapWithCapacity[K comparable, V any](cap uint) *HashMap[K, V] {
	return NewHashMapWith[K, V](cap, HashDefaultFunc[K])
}
func NewHashMapWithHasher[K, V any](hasher func(K) uint64) *HashMap[K, V] {
	return NewHashMapWith[K, V](defaultInitCapacity, hasher)
}
func NewHashMapWith[K, V any](cap uint, hasher func(K) uint64) *HashMap[K, V] {
	hm := &HashMap[K, V]{
		hasher: hasher,
		data:   make([]entry[K, V], cap),
	}
	return hm
}

func (hm *HashMap[K, V]) hash(k K) uint64 {
	return hm.hasher(k)
}
func (hm *HashMap[K, V]) index(hash uint64) uint {
	return uint(hash & uint64(len(hm.data)-1))
}
func (hm *HashMap[K, V]) expan() {
	oldData := hm.data
	hm.data = make([]entry[K, V], len(oldData)*2)
	for _, entry := range oldData {
		for cursor := entry.next; cursor != nil; {
			next := cursor.next

			prev, _ := hm.find(cursor.hash, cursor.key)
			prev.next = cursor
			cursor.next = nil

			cursor = next
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
func (hm *HashMap[K, V]) find(hash uint64, k K) (prev *entry[K, V], cursor *entry[K, V]) {
	entry := &hm.data[hm.index(hash)]
	for prev, cursor = entry, entry.next; cursor != nil; prev, cursor = cursor, cursor.next {
		if hash == cursor.hash {
			return prev, cursor
		}
	}
	return prev, nil
}
func (hm *HashMap[K, V]) Set(k K, v V) bool {
	hm.checkExpan()

	hash := hm.hash(k)

	prev, cursor := hm.find(hash, k)
	if cursor != nil {
		cursor.value = v
		return false
	}

	prev.next = newEntry(hash, k, v)
	hm.len++
	return true
}
func (hm *HashMap[K, V]) Get(k K, v ...V) (V, bool) {
	hash := hm.hash(k)

	_, cursor := hm.find(hash, k)
	if cursor != nil {
		return cursor.value, true
	}

	if len(v) > 0 {
		return v[len(v)-1], false
	}

	var val V
	return val, false
}
func (hm *HashMap[K, V]) ContainKey(k K) bool {
	hash := hm.hash(k)
	_, cursor := hm.find(hash, k)
	return cursor != nil
}
func (hm *HashMap[K, V]) Remove(k K, v ...V) (V, bool) {
	hash := hm.hash(k)

	prev, cursor := hm.find(hash, k)
	if cursor != nil {
		hm.len--
		prev.next = cursor.next
		return cursor.value, true
	}

	if len(v) > 0 {
		return v[len(v)-1], false
	}

	var val V
	return val, false
}
func (hm *HashMap[K, V]) Clear() {
	hm.data = make([]entry[K, V], defaultInitCapacity)
	hm.len = 0
}
func (hm *HashMap[K, V]) Keys() []K {
	if hm.len == 0 {
		return nil
	}

	keys := make([]K, hm.Length())
	var i int
	for _, entry := range hm.data {
		for cursor := entry.next; cursor != nil; cursor = cursor.next {
			keys[i] = cursor.key
			i++
		}
	}
	rand.Shuffle(
		len(keys), func(i, j int) {
			keys[i], keys[j] = keys[j], keys[i]
		},
	)

	return keys
}
func (hm *HashMap[K, V]) Values() []V {
	if hm.len == 0 {
		return nil
	}

	vals := make([]V, hm.Length())
	var i int
	for _, entry := range hm.data {
		for cursor := entry.next; cursor != nil; cursor = cursor.next {
			vals[i] = cursor.value
			i++
		}
	}
	rand.Shuffle(
		len(vals), func(i, j int) {
			vals[i], vals[j] = vals[j], vals[i]
		},
	)

	return vals
}
