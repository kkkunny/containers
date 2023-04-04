package iter

type Iter[T any] interface {
	Value() T
	HasNext() bool
	Next()
}
