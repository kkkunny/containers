package linkedhashmap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/kkkunny/containers/hashmap"
)

func TestNewLinkedHashMap(t *testing.T) {
	hm := NewLinkedHashMap[int, int]()
	if hm.Length() != 0 {
		t.FailNow()
	}
}

func TestLinkedHashMap(t *testing.T) {
	hm := NewLinkedHashMap[int, int]()
	for i := 0; i < 10; i++ {
		hm.Set(i, i)
	}
	if hm.Length() != 10 {
		t.FailNow()
	}
	if v, ok := hm.Get(11, 12); ok || v != 12 {
		t.FailNow()
	}
	if v, ok := hm.Get(0, 13); !ok || v != 0 {
		t.FailNow()
	}
	if !hm.ContainKey(1) {
		t.FailNow()
	}
	if _, v := hm.Back(); v != 9 {
		t.FailNow()
	}
}

func BenchmarkMix_HashMap(b *testing.B) {
	rand.Seed(time.Now().Unix())
	hm := NewLinkedHashMapWithHasher[int, int](hashmap.HashSignedIntFunc[int])
	for i := 1; i <= 1000; i++ {
		hm.Set(i, i)
		for j := 0; j < 10; j++ {
			_, _ = hm.Get(rand.Intn(i))
		}
	}
}

func BenchmarkMix_Stdlib(b *testing.B) {
	rand.Seed(time.Now().Unix())
	hm := hashmap.NewHashMapWithHasher[int, int](hashmap.HashSignedIntFunc[int])
	for i := 1; i <= 1000; i++ {
		hm.Set(i, i)
		for j := 0; j < 10; j++ {
			_, _ = hm.Get(rand.Intn(i))
		}
	}
}
