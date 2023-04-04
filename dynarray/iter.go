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
func (i *DynArrayIter[T]) Prev() *DynArrayIter[T] {
	if !i.HasPrev() {
		return nil
	}
	return &DynArrayIter[T]{
		da:    i.da,
		index: i.index - 1,
	}
}
func (i *DynArrayIter[T]) HasNext() bool {
	return i.index != uint(len(i.da.data))-1
}
func (i *DynArrayIter[T]) Next() *DynArrayIter[T] {
	if !i.HasNext() {
		return nil
	}
	return &DynArrayIter[T]{
		da:    i.da,
		index: i.index + 1,
	}
}
