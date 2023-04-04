package dynarray

type DynArray[T any] struct {
	data []T
}

func NewDynArray[T any]() *DynArray[T] {
	return &DynArray[T]{data: make([]T, 0, 0)}
}
func NewDynArrayWith[T any](len, cap uint) *DynArray[T] {
	return &DynArray[T]{data: make([]T, len, cap)}
}

func (da *DynArray[T]) checkIndex(i uint) {
	if i > uint(len(da.data)) {
		panic("index out of range")
	}
}
func (da *DynArray[T]) Length() uint {
	return uint(len(da.data))
}
func (da *DynArray[T]) Capacity() uint {
	return uint(cap(da.data))
}
func (da *DynArray[T]) Add(v T) T {
	da.data = append(da.data, v)
	return v
}
func (da *DynArray[T]) Get(i uint) T {
	return da.data[i]
}
func (da *DynArray[T]) Insert(i uint, v T) T {
	da.checkIndex(i)
	da.data = append(da.data, v)
	copy(da.data[i+1:], da.data[i:])
	da.data[i] = v
	return v
}
func (da *DynArray[T]) Slice(begin, end uint) *DynArray[T] {
	length := end - begin
	if length <= 0 || begin >= uint(len(da.data)) {
		return NewDynArray[T]()
	}

	newDa := NewDynArrayWith[T](length, length)
	if end > uint(len(da.data)) {
		copy(newDa.data, da.data[begin:])
	} else {
		copy(newDa.data, da.data[begin:end])
	}
	return newDa
}
func (da *DynArray[T]) Set(i uint, v T) T {
	da.checkIndex(i)
	da.data[i] = v
	return v
}
func (da *DynArray[T]) Remove(i uint) T {
	da.checkIndex(i)

	v := da.data[i]

	if i == 0 {
		da.data = da.data[1:]
	} else if i == uint(len(da.data))-1 {
		da.data = da.data[:len(da.data)-1]
	} else {
		copy(da.data[i:], da.data[i+1:])
		da.data = da.data[:len(da.data)-1]
	}

	return v
}
