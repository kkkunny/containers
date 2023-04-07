package bstree

import (
	"testing"
)

func TestNewBSTree(t *testing.T) {
	bst := NewBSTree[int]()
	if bst.Length() != 0 || bst.Deepth() != 0 {
		t.FailNow()
	}
}

func TestBSTree(t *testing.T) {
	bst := NewBSTree[int]()
	for i := 0; i < 15; i++ {
		bst.Push(i)
	}
	if bst.Length() != 15 {
		t.FailNow()
	}
	if !bst.Contain(1) {
		t.FailNow()
	}
	if v, ok := bst.Get(1); !ok || v != 1 {
		t.FailNow()
	}
	if v, ok := bst.Remove(1, 2); !ok || v != 1 {
		t.FailNow()
	}
	if v, ok := bst.Get(1, 2); ok || v != 2 {
		t.FailNow()
	}
}
