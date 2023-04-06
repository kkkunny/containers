package hashmap

import (
	"fmt"
	"testing"
)

func TestHashMap_Iter(t *testing.T) {
	// TODO: failed
	hm := NewHashMap[string, int]()
	for i := 0; i <= 2; i++ {
		hm.Set(fmt.Sprintf("%d", i), i)
	}
	fmt.Println(hm)
	for iter := hm.Iterator(); iter != nil; iter.Next() {
		fmt.Println(iter.Key(), fmt.Sprintf("%v", iter.Value()))
		// if iter.Key() != fmt.Sprintf("%v", iter.Value()) {
		// 	t.FailNow()
		// }
		if !iter.HasNext() {
			break
		}
	}
	t.Fail()
}
