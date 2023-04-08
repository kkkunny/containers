package hashmap

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewHashMap(t *testing.T) {
	hm := NewHashMapWithHasher[int, int](HashSignedIntFunc[int])
	if hm.Length() != 0 {
		t.FailNow()
	}
}

func TestHashMap(t *testing.T) {
	hm := NewHashMapWithHasher[int, int](HashSignedIntFunc[int])
	for i := 0; i < 15; i++ {
		hm.Set(i, i)
	}
	if hm.Length() != 15 {
		t.FailNow()
	}
	if v, ok := hm.Get(16, 12); ok || v != 12 {
		t.FailNow()
	}
	if v, ok := hm.Get(0, 13); !ok || v != 0 {
		t.FailNow()
	}
	if v, ok := hm.Remove(0, 1); !ok || v != 0 {
		t.FailNow()
	}
	if v, ok := hm.Get(0, 13); ok || v != 13 {
		t.FailNow()
	}
	if !hm.ContainKey(1) {
		t.FailNow()
	}
}

func BenchmarkAdd_HashMap(b *testing.B) {
	hm := NewHashMapWithHasher[int, int](HashSignedIntFunc[int])
	for i := 0; i < 100000; i++ {
		hm.Set(i, i)
	}
}

func BenchmarkAdd_Stdlib(b *testing.B) {
	hm := make(map[int]int)
	for i := 0; i < 100000; i++ {
		hm[i] = i
	}
}

func BenchmarkGet_HashMap(b *testing.B) {
	rand.Seed(time.Now().Unix())
	hm := NewHashMapWithHasher[int, int](HashSignedIntFunc[int])
	for i := 0; i < 100; i++ {
		hm.Set(i, i)
	}
	for i := 0; i < 10000; i++ {
		_, _ = hm.Get(rand.Intn(100))
	}
}

func BenchmarkGet_Stdlib(b *testing.B) {
	rand.Seed(time.Now().Unix())
	hm := make(map[int]int)
	for i := 0; i < 100; i++ {
		hm[i] = i
	}
	for i := 0; i < 10000; i++ {
		_ = hm[rand.Intn(100)]
	}
}

func BenchmarkMix_HashMap(b *testing.B) {
	rand.Seed(time.Now().Unix())
	hm := NewHashMapWithHasher[int, int](HashSignedIntFunc[int])
	for i := 1; i <= 1000; i++ {
		hm.Set(i, i)
		for j := 0; j < 10; j++ {
			_, _ = hm.Get(rand.Intn(i))
		}
	}
}

func BenchmarkMix_Stdlib(b *testing.B) {
	rand.Seed(time.Now().Unix())
	hm := make(map[int]int)
	for i := 1; i <= 1000; i++ {
		hm[i] = i
		for j := 0; j < 10; j++ {
			_ = hm[rand.Intn(i)]
		}
	}
}
