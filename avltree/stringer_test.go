package avltree

import "testing"

func TestAVLTree_String(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Push(5)
	tree.Push(1)
	tree.Push(7)
	tree.Push(3)
	tree.Push(9)
	tree.Push(0)
	tree.Push(2)
	tree.Push(4)
	tree.Push(6)
	tree.Push(8)
	if tree.String() != "{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}" {
		t.FailNow()
	}
}
