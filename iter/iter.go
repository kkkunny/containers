package iter

type Iter[T any] interface {
	Value() T
	HasPrev() bool
	HasNext() bool
}

type IndexIter[T any] interface {
	Iter[T]
	Index() uint
}
