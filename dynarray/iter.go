package dynarray

type DynArrayIter[T any] struct {
	da    *DynArray[T]
	index uint
}

func (da *DynArray[T]) Begin() *DynArrayIter[T] {
	if len(da.data) == 0 {
		return nil
	}
	return &DynArrayIter[T]{
		da:    da,
		index: 0,
	}
}
func (da *DynArray[T]) End() *DynArrayIter[T] {
	if len(da.data) == 0 {
		return nil
	}
	return &DynArrayIter[T]{
		da:    da,
		index: uint(len(da.data) - 1),
	}
}

func (i *DynArrayIter[T]) Index() uint {
	return i.index
}
func (i *DynArrayIter[T]) Value() T {
	return i.da.data[i.index]
}
func (i *DynArrayIter[T]) HasPrev() bool {
	return i.index != 0
}
func (i *DynArrayIter[T]) Prev() {
	if !i.HasPrev() {
		panic("index out of range")
	}
	i.index--
}
func (i *DynArrayIter[T]) HasNext() bool {
	return i.index != uint(len(i.da.data))-1
}
func (i *DynArrayIter[T]) Next() {
	if !i.HasNext() {
		panic("index out of range")
	}
	i.index++
}

func DynArrayMap[From, To any](da *DynArray[From], fn func(i uint, v From)To)*DynArray[To]{
	newDa := NewDynArrayWith[To](da.Length(), da.Capacity())
	for iter:=da.Begin(); iter != nil; iter.Next() {
		newDa.Set(iter.Index(), fn(iter.Index(), iter.Value()))
		if !iter.HasNext(){
			break
		}
	}
	return newDa
}
