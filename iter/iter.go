package iter

type Iter[T any] interface {
	Value() T
	HasPrev() bool
	Prev()
	HasNext() bool
	Next()
}
