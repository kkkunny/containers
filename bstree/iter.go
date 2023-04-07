package bstree

type BSTreeIter[T any] struct {
	tree   *BSTree[T]
	cursor *node[T]
}

func (tree *BSTree[T]) Begin() *BSTreeIter[T] {
	if tree.root == nil {
		return nil
	}
	return &BSTreeIter[T]{
		tree:   tree,
		cursor: tree.root.leftmost(),
	}
}
func (tree *BSTree[T]) End() *BSTreeIter[T] {
	if tree.root == nil {
		return nil
	}
	return &BSTreeIter[T]{
		tree:   tree,
		cursor: tree.root.rightmost(),
	}
}

func (i *BSTreeIter[T]) Value() T {
	return i.cursor.value
}
func (i *BSTreeIter[T]) HasPrev() bool {
	yln := i.cursor.youngerLeftNeighbor()
	if yln != nil {
		return true
	}
	return i.cursor.olderLeftNeighbor(i.tree.order) != nil
}
func (i *BSTreeIter[T]) Prev() {
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
func (i *BSTreeIter[T]) HasNext() bool {
	yrn := i.cursor.youngerRightNeighbor()
	if yrn != nil {
		return true
	}
	return i.cursor.olderRightNeighbor(i.tree.order) != nil
}
func (i *BSTreeIter[T]) Next() {
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
