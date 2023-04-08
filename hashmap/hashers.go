package hashmap

import (
	"hash/fnv"
	"math"
	"reflect"
	"unsafe"

	"golang.org/x/exp/constraints"
)

func HashDefaultFunc[T any](v T) uint64 {
	h := fnv.New64a()
	size := reflect.TypeOf(v).Size()
	ptr := uintptr(unsafe.Pointer(&v))
	for i := uintptr(0); i < size; i++ {
		b := *(*byte)(unsafe.Pointer(ptr + i))
		_, _ = h.Write([]byte{b})
	}
	return h.Sum64()
}

func HashBoolFunc[T ~bool](v T) uint64 {
	if v {
		return 1
	}
	return 0
}

func HashStringFunc[T ~string](v T) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(v))
	return h.Sum64()
}

func HashSignedIntFunc[T constraints.Signed](v T) uint64 {
	int64v := int64(v)
	uint64v := *(*uint64)(unsafe.Pointer(&int64v))
	return uint64v
}

func HashUnsignedIntFunc[T constraints.Unsigned](v T) uint64 {
	return uint64(v)
}

func HashFloatFunc[T constraints.Float](v T) uint64 {
	return math.Float64bits(float64(v))
}
