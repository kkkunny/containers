package list

import "testing"

func TestList_Iter(t *testing.T) {
	l := NewList[int]()
	for i := 0; i <= 2; i++ {
		l.PushBack(i)
	}
	var i int
	for iter := l.Begin(); iter != nil; iter.Next() {
		if i != iter.Value() {
			t.FailNow()
		}
		i++
		if !iter.HasNext(){
			break
		}
	}
}
