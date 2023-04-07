package bstree

import "testing"

func TestDynArray_String(t *testing.T) {
	bst := NewBSTree[int]()
	bst.Push(5)
	bst.Push(1)
	bst.Push(7)
	bst.Push(3)
	bst.Push(9)
	bst.Push(0)
	bst.Push(2)
	bst.Push(4)
	bst.Push(6)
	bst.Push(8)
	if bst.String() != "{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}" {
		t.FailNow()
	}
}
