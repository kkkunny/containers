package avltree

import (
	"testing"
)

func TestNewAVLTree(t *testing.T) {
	tree := NewAVLTree[int]()
	if tree.Length() != 0 || tree.Deepth() != 0 {
		t.FailNow()
	}
}

func TestAVLTree(t *testing.T) {
	tree := NewAVLTree[int]()
	for i := 0; i < 15; i++ {
		tree.Push(i)
	}
	if tree.Length() != 15 {
		t.FailNow()
	}
	if !tree.Contain(1) {
		t.FailNow()
	}
	if v, ok := tree.Get(1); !ok || v != 1 {
		t.FailNow()
	}
	if v, ok := tree.Remove(1, 2); !ok || v != 1 {
		t.FailNow()
	}
	if v, ok := tree.Get(1, 2); ok || v != 2 {
		t.FailNow()
	}
}
