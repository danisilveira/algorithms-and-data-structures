package binarysearch

import "cmp"

type preorderIterator[T cmp.Ordered] struct {
	tree     *Tree[T]
	node     *TreeNode[T]
	lastNode *TreeNode[T]
	finished bool
}

func NewPreOrderIterator[T cmp.Ordered](tree *Tree[T]) TreeIterator[T] {
	return &preorderIterator[T]{
		tree:     tree,
		lastNode: tree.Right(tree.root),
	}
}

func (i *preorderIterator[T]) HasNext() bool {
	if i.finished {
		return false
	}

	if i.node == nil {
		i.node = i.tree.root
		return i.node != nil
	}

	if i.node.Left != nil {
		i.node = i.node.Left
		return true
	}

	if i.node.Right != nil {
		i.node = i.node.Right
		return true
	}

	for i.node != i.lastNode {
		node := i.node
		i.node = i.node.Parent

		if i.node.Right != nil && node != i.node.Right {
			i.node = i.node.Right
			return true
		}
	}

	i.finished = true
	return false
}

func (i *preorderIterator[T]) Value() T {
	return i.node.Value
}

func (i *preorderIterator[T]) Reset() {
	i.node = nil
	i.finished = false
}
