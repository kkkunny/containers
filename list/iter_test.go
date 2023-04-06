package list

import (
	"testing"
)

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
		if !iter.HasNext() {
			break
		}
	}
}

func TestList_Map(t *testing.T) {
	l := NewList[int]()
	for i := 0; i <= 2; i++ {
		l.PushBack(i)
	}
	nl := ListMap(l, func(v int) int {
		return v + 1
	})
	var i int
	for iter := nl.Begin(); iter != nil; iter.Next() {
		if i != iter.Value()-1 {
			t.FailNow()
		}
		i++
		if !iter.HasNext() {
			break
		}
	}
}
