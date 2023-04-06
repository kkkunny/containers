package hashmap

import (
	"testing"
)

func TestNewHashMap(t *testing.T) {
	hm := NewHashMap[string, int]()
	if hm.Length() != 0 {
		t.FailNow()
	}
}

func TestHashMap(t *testing.T) {
	hm := NewHashMap[int, int]()
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
	if !hm.ContainKey(1) {
		t.FailNow()
	}
}
