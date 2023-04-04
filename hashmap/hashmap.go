package hashmap

type HashMap[K comparable, V any] struct{
	data map[K]V
}

func NewHashMap[K comparable, V any]()*HashMap[K, V]{
	return &HashMap[K, V]{data: make(map[K]V)}
}
func NewHashMapWith[K comparable, V any](cap uint)*HashMap[K, V]{
	return &HashMap[K, V]{data: make(map[K]V, cap)}
}

func (hm *HashMap[K, V]) Length() uint {
	return uint(len(hm.data))
}
func (hm *HashMap[K, V]) Set(k K, v V) V {
	hm.data[k] = v
	return v
}
func (hm *HashMap[K, V]) Get(k K, v ...V) V {
	val, ok := hm.data[k]
	if !ok && len(v) > 0{
		return v[len(v)-1]
	}
	return val
}
func (hm *HashMap[K, V]) ContainKey(k K) bool {
	_, ok := hm.data[k]
	return ok
}
func (hm *HashMap[K, V]) Remove(k K, v ...V) V {
	val, ok := hm.data[k]
	if !ok{
		if len(v) != 0{
			return v[len(v)-1]
		}
		return val
	}
	delete(hm.data, k)
	return val
}
