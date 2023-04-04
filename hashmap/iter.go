package hashmap

type HashMapIter[K comparable, V any] struct {
	hm *HashMap[K, V]
	keys []K
	index uint
}

func (hm *HashMap[K, V]) Iterator() *HashMapIter[K, V] {
	if len(hm.data) == 0 {
		return nil
	}
	keys := make([]K, hm.Length(), hm.Length())
	var i int
	for k := range hm.data{
		keys[i] = k
		i++
	}
	return &HashMapIter[K, V]{
		hm:    hm,
		keys: keys,
		index: 0,
	}
}

func (i *HashMapIter[K, V]) Key() K {
	return i.keys[i.index]
}
func (i *HashMapIter[K, V]) Value() V {
	return i.hm.Get(i.Key())
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

func HashMapMap[FK comparable, FV any, TK comparable, TV any](hm *HashMap[FK, FV], fn func(k FK, v FV)(TK, TV))*HashMap[TK, TV]{
	newHm := NewHashMapWith[TK, TV](hm.Length())
	for iter:=hm.Iterator(); iter != nil; iter.Next() {
		newHm.Set(fn(iter.Key(), iter.Value()))
		if !iter.HasNext(){
			break
		}
	}
	return newHm
}
