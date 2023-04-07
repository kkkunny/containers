package bstree

import (
	"golang.org/x/exp/constraints"
)

type node[T any] struct {
	father, left, right *node[T]
	value               T
}

func (n *node[T]) deepth() uint {
	if n.left == nil && n.right == nil {
		return 1
	} else if n.left == nil {
		return n.right.deepth() + 1
	} else if n.right == nil {
		return n.left.deepth() + 1
	} else {
		l, r := n.left.deepth(), n.right.deepth()
		if l < r {
			return r + 1
		} else {
			return l + 1
		}
	}
}
func (n *node[T]) setLeft(o *node[T]) {
	n.left = o
	if o != nil {
		o.father = n
	}
}
func (n *node[T]) setRight(o *node[T]) {
	n.right = o
	if o != nil {
		o.father = n
	}
}
func (n *node[T]) setFatherThisSon(o *node[T]) bool {
	if n.father == nil {
		return false
	}
	if n.father.left == n {
		n.father.left = o
		if o != nil {
			o.father = n.father
		}
		return true
	} else if n.father.right == n {
		n.father.right = o
		if o != nil {
			o.father = n.father
		}
		return true
	} else {
		return false
	}
}
func (n *node[T]) leftmost() *node[T] {
	if n.left == nil {
		return n
	}
	return n.left.leftmost()
}
func (n *node[T]) rightmost() *node[T] {
	if n.right == nil {
		return n
	}
	return n.right.rightmost()
}
func (n *node[T]) youngerLeftNeighbor() *node[T] {
	if n.left == nil {
		return nil
	}
	return n.left.rightmost()
}
func (n *node[T]) youngerRightNeighbor() *node[T] {
	if n.right == nil {
		return nil
	}
	return n.right.leftmost()
}
func (n *node[T]) olderLeftNeighbor(order func(T, T) int) *node[T] {
	me, father := n, n.father
	if father == nil {
		return nil
	}

	if order(me.value, father.value) > 0 {
		return father
	}

	return father.olderLeftNeighbor(order)
}
func (n *node[T]) olderRightNeighbor(order func(T, T) int) *node[T] {
	me, father := n, n.father
	if father == nil {
		return nil
	}

	if order(me.value, father.value) < 0 {
		return father
	}

	return father.olderRightNeighbor(order)
}

type BSTree[T any] struct {
	order func(T, T) int
	root  *node[T]
	len   uint
}

func NewBSTree[T constraints.Ordered]() *BSTree[T] {
	return &BSTree[T]{order: func(l, r T) int {
		if l < r {
			return -1
		} else if l == r {
			return 0
		} else {
			return 1
		}
	}}
}
func NewBSTreeWith[T any](order func(T, T) int) *BSTree[T] {
	return &BSTree[T]{order: order}
}

func (tree *BSTree[T]) Length() uint {
	return tree.len
}
func (tree *BSTree[T]) Deepth() uint {
	if tree.root == nil {
		return 0
	}
	return tree.root.deepth()
}
func (tree *BSTree[T]) findNode(cursor *node[T], v T) (*node[T], *node[T]) {
	if res := tree.order(v, cursor.value); res < 0 {
		if cursor.left == nil {
			return cursor, nil
		} else {
			return tree.findNode(cursor.left, v)
		}
	} else if res == 0 {
		return cursor.father, cursor
	} else {
		if cursor.right == nil {
			return cursor, nil
		} else {
			return tree.findNode(cursor.right, v)
		}
	}
}
func (tree *BSTree[T]) Push(v T) bool {
	newNode := &node[T]{value: v}

	if tree.root == nil {
		tree.root = newNode
		tree.len++
		return true
	}

	father, node := tree.findNode(tree.root, v)
	if node != nil {
		node.value = v
		return false
	}

	if res := tree.order(v, father.value); res < 0 {
		father.setLeft(newNode)
	} else {
		father.setRight(newNode)
	}
	tree.len++
	return true
}
func (tree *BSTree[T]) Peek() T {
	return tree.root.value
}
func (tree *BSTree[T]) removeNode(n *node[T]) *node[T] {
	if n.left == nil && n.right == nil {
		n.setFatherThisSon(nil)
	} else if n.left == nil {
		newNode := tree.removeNode(n.right.leftmost())
		tree.len++
		newNode.setRight(n.right)
		if !n.setFatherThisSon(newNode) {
			tree.root = newNode
		}
	} else if n.right == nil {
		newNode := tree.removeNode(n.left.rightmost())
		tree.len++
		newNode.setLeft(n.left)
		if !n.setFatherThisSon(newNode) {
			tree.root = newNode
		}
	} else {
		ld, rd := n.left.deepth(), n.right.deepth()
		var newNode *node[T]
		if ld < rd {
			newNode = tree.removeNode(n.right.leftmost())
		} else {
			newNode = tree.removeNode(n.left.rightmost())
		}
		tree.len++
		newNode.setLeft(n.left)
		newNode.setRight(n.right)
		if !n.setFatherThisSon(newNode) {
			tree.root = newNode
		}
	}

	if tree.root == n {
		tree.root = nil
	}
	tree.len--
	return n
}
func (tree *BSTree[T]) Pop() T {
	return tree.removeNode(tree.root).value
}
func (tree *BSTree[T]) Remove(v T, defaultValue ...T) (T, bool) {
	if tree.root == nil && len(defaultValue) > 0 {
		return defaultValue[len(defaultValue)-1], false
	} else if tree.root == nil {
		var tmp T
		return tmp, false
	}

	_, node := tree.findNode(tree.root, v)
	if node == nil && len(defaultValue) > 0 {
		return defaultValue[len(defaultValue)-1], false
	} else if node == nil {
		var tmp T
		return tmp, false
	}
	return tree.removeNode(node).value, true
}
func (tree *BSTree[T]) Contain(v T) bool {
	if tree.root == nil {
		return false
	}

	_, node := tree.findNode(tree.root, v)
	return node != nil
}
func (tree *BSTree[T]) Clear() {
	tree.root = nil
	tree.len = 0
}
func (tree *BSTree[T]) Get(v T, defaultValue ...T) (T, bool) {
	if tree.root == nil && len(defaultValue) > 0 {
		return defaultValue[len(defaultValue)-1], false
	} else if tree.root == nil {
		var tmp T
		return tmp, false
	}

	_, node := tree.findNode(tree.root, v)
	if node == nil && len(defaultValue) > 0 {
		return defaultValue[len(defaultValue)-1], false
	} else if node == nil {
		var tmp T
		return tmp, false
	}
	return node.value, true
}
