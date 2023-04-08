package avltree

type AVLTreeIter[T any] struct {
	tree   *AVLTree[T]
	cursor *node[T]
}

func (tree *AVLTree[T]) Begin() *AVLTreeIter[T] {
	if tree.root == nil {
		return nil
	}
	return &AVLTreeIter[T]{
		tree:   tree,
		cursor: tree.root.leftmost(),
	}
}
func (tree *AVLTree[T]) End() *AVLTreeIter[T] {
	if tree.root == nil {
		return nil
	}
	return &AVLTreeIter[T]{
		tree:   tree,
		cursor: tree.root.rightmost(),
	}
}

func (i *AVLTreeIter[T]) Value() T {
	return i.cursor.value
}
func (i *AVLTreeIter[T]) HasPrev() bool {
	yln := i.cursor.youngerLeftNeighbor()
	if yln != nil {
		return true
	}
	return i.cursor.olderLeftNeighbor(i.tree.order) != nil
}
func (i *AVLTreeIter[T]) Prev() {
	yln := i.cursor.youngerLeftNeighbor()
	if yln != nil {
		i.cursor = yln
		return
	}
	oln := i.cursor.olderLeftNeighbor(i.tree.order)
	if oln != nil {
		i.cursor = oln
		return
	}
	panic("index out of range")
}
func (i *AVLTreeIter[T]) HasNext() bool {
	yrn := i.cursor.youngerRightNeighbor()
	if yrn != nil {
		return true
	}
	return i.cursor.olderRightNeighbor(i.tree.order) != nil
}
func (i *AVLTreeIter[T]) Next() {
	yrn := i.cursor.youngerRightNeighbor()
	if yrn != nil {
		i.cursor = yrn
		return
	}
	orn := i.cursor.olderRightNeighbor(i.tree.order)
	if orn != nil {
		i.cursor = orn
		return
	}
	panic("index out of range")
}
