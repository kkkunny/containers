package treemap

import (
	"github.com/kkkunny/containers/bstree"
	"golang.org/x/exp/constraints"
)

type entry[K, V any] struct {
	key   K
	value V
}

type TreeMap[K, V any] struct {
	tree *bstree.BSTree[entry[K, V]]
}

func NewTreeMap[K constraints.Ordered, V any]() *TreeMap[K, V] {
	return NewTreeMapWith[K, V](func(l, r K) int {
		if l < r {
			return -1
		} else if l == r {
			return 0
		} else {
			return 1
		}
	})
}
func NewTreeMapWith[K, V any](order func(K, K) int) *TreeMap[K, V] {
	return &TreeMap[K, V]{tree: bstree.NewBSTreeWith[entry[K, V]](func(l, r entry[K, V]) int {
		return order(l.key, r.key)
	})}
}

func (m *TreeMap[K, V]) Length() uint {
	return m.tree.Length()
}
func (m *TreeMap[K, V]) Set(k K, v V) bool {
	return m.tree.Push(entry[K, V]{
		key:   k,
		value: v,
	})
}
func (m *TreeMap[K, V]) Get(k K, defaultValue ...V) (V, bool) {
	v, ok := m.tree.Get(entry[K, V]{key: k})
	if !ok && len(defaultValue) > 0 {
		return defaultValue[len(defaultValue)-1], false
	} else if !ok {
		var tmp V
		return tmp, false
	}
	return v.value, true
}
func (m *TreeMap[K, V]) ContainKey(k K) bool {
	return m.tree.Contain(entry[K, V]{key: k})
}
func (m *TreeMap[K, V]) Remove(k K, defaultValue ...V) (V, bool) {
	v, ok := m.tree.Remove(entry[K, V]{key: k})
	if !ok && len(defaultValue) > 0 {
		return defaultValue[len(defaultValue)-1], false
	} else if !ok {
		var tmp V
		return tmp, false
	}
	return v.value, true
}
func (m *TreeMap[K, V]) Clear() {
	m.tree.Clear()
}
func (m *TreeMap[K, V]) Keys() []K {
	if m.tree.Length() == 0 {
		return nil
	}

	keys := make([]K, m.tree.Length())
	var i int
	for iter := m.tree.Begin(); iter != nil; iter.Next() {
		keys[i] = iter.Value().key
		if !iter.HasNext() {
			break
		}
		i++
	}
	return keys
}
func (m *TreeMap[K, V]) Values() []V {
	if m.tree.Length() == 0 {
		return nil
	}

	vals := make([]V, m.tree.Length())
	var i int
	for iter := m.tree.Begin(); iter != nil; iter.Next() {
		vals[i] = iter.Value().value
		if !iter.HasNext() {
			break
		}
		i++
	}
	return vals
}
