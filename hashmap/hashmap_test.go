package hashmap

import "testing"

func TestNewHashMap(t *testing.T) {
	hm := NewHashMapWith[int, int](10)
	if hm.Length() != 0{
		t.FailNow()
	}
}

func TestHashMap(t *testing.T){
	hm := NewHashMapWith[int, int](10)
	for i:=0; i <10; i++{
		hm.Set(i, i)
	}
	if hm.Length() != 10{
		t.FailNow()
	}
	if hm.Get(11, 12) != 12{
		t.FailNow()
	}
	if hm.Get(0, 13) != 0{
		t.FailNow()
	}
	if !hm.ContainKey(1){
		t.FailNow()
	}
}
