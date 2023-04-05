package linkedhashmap

import "testing"

func TestNewLinkedHashMap(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](10)
	if hm.Length() != 0 {
		t.FailNow()
	}
}

func TestLinkedHashMap(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](10)
	for i := 0; i < 10; i++ {
		hm.Set(i, i)
	}
	if hm.Length() != 10 {
		t.FailNow()
	}
	if hm.Get(11, 12) != 12 {
		t.FailNow()
	}
	if hm.Get(0, 13) != 0 {
		t.FailNow()
	}
	if !hm.ContainKey(1) {
		t.FailNow()
	}
	if _, v := hm.Back(); v != 9 {
		t.FailNow()
	}
}
