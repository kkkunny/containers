package hashmap

type HashMapIter[K, V any] struct {
	hm    *HashMap[K, V]
	keys  []K
	index uint
}

func (hm *HashMap[K, V]) Iterator() *HashMapIter[K, V] {
	if hm.len == 0 {
		return nil
	}
	return &HashMapIter[K, V]{
		hm:    hm,
		keys:  hm.Keys(),
		index: 0,
	}
}

func (i *HashMapIter[K, V]) Key() K {
	return i.keys[i.index]
}
func (i *HashMapIter[K, V]) Value() V {
	v, _ := i.hm.Get(i.Key())
	return v
}
func (i *HashMapIter[K, V]) HasNext() bool {
	return i.index != uint(len(i.keys))-1
}
func (i *HashMapIter[K, V]) Next() {
	if !i.HasNext() {
		panic("index out of range")
	}
	i.index++
}
