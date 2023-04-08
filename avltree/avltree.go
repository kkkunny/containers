package avltree

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
	if o != nil {
		o.father = n.father
	}
	if n.father == nil {
		return false
	}
	if n.father.left == n {
		n.father.left = o
		return true
	} else if n.father.right == n {
		n.father.right = o
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
func (n *node[T]) balanceFactor() int {
	var ld, rd int
	if n.left != nil {
		ld = int(n.left.deepth())
	}
	if n.right != nil {
		rd = int(n.right.deepth())
	}
	return ld - rd
}
func (n *node[T]) leftRotate() *node[T] {
	root := n.right
	n.setFatherThisSon(root)
	n.setRight(root.left)
	root.setLeft(n)
	return root
}
func (n *node[T]) rightRotate() *node[T] {
	root := n.left
	n.setFatherThisSon(root)
	n.setLeft(root.right)
	root.setRight(n)
	return root
}

type AVLTree[T any] struct {
	order func(T, T) int
	root  *node[T]
	len   uint
}

func NewAVLTree[T constraints.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{
		order: func(l, r T) int {
			if l < r {
				return -1
			} else if l == r {
				return 0
			} else {
				return 1
			}
		},
	}
}
func NewAVLTreeWith[T any](order func(T, T) int) *AVLTree[T] {
	return &AVLTree[T]{order: order}
}

func (tree *AVLTree[T]) Length() uint {
	return tree.len
}
func (tree *AVLTree[T]) Deepth() uint {
	if tree.root == nil {
		return 0
	}
	return tree.root.deepth()
}
func (tree *AVLTree[T]) findNode(cursor *node[T], v T) (*node[T], *node[T]) {
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
func (tree *AVLTree[T]) Push(v T) bool {
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

	tree.checkBalance(newNode, true)
	return true
}
func (tree *AVLTree[T]) Peek() T {
	return tree.root.value
}
func (tree *AVLTree[T]) removeNode(n *node[T]) *node[T] {
	if n.left == nil && n.right == nil {
		if !n.setFatherThisSon(nil) {
			tree.root = n
		}
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
		var newNode *node[T]
		if balance := n.balanceFactor(); balance <= 0 {
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

	tree.len--
	if n.father != nil {
		tree.checkBalance(n.father, false)
	}
	return n
}
func (tree *AVLTree[T]) Pop() T {
	return tree.removeNode(tree.root).value
}
func (tree *AVLTree[T]) Remove(v T, defaultValue ...T) (T, bool) {
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
func (tree *AVLTree[T]) Contain(v T) bool {
	if tree.root == nil {
		return false
	}

	_, node := tree.findNode(tree.root, v)
	return node != nil
}
func (tree *AVLTree[T]) Clear() {
	tree.root = nil
	tree.len = 0
}
func (tree *AVLTree[T]) Get(v T, defaultValue ...T) (T, bool) {
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
func (tree *AVLTree[T]) checkBalance(cursor *node[T], onlyOne bool) {
	change := true
	if balence := cursor.balanceFactor(); balence < -1 {
		if cursor.right.right != nil {
			cursor = cursor.leftRotate()
		} else if cursor.right.left != nil {
			cursor.left.leftRotate()
			cursor = cursor.rightRotate()
		} else {
			panic("unreachable")
		}
	} else if balence > 1 {
		if cursor.left.left != nil {
			cursor = cursor.rightRotate()
		} else if cursor.left.right != nil {
			cursor.right.rightRotate()
			cursor = cursor.leftRotate()
		} else {
			panic("unreachable")
		}
	} else {
		change = false
	}
	if cursor.father == nil {
		tree.root = cursor
	} else if !onlyOne || !change {
		tree.checkBalance(cursor.father, change)
	}
}
