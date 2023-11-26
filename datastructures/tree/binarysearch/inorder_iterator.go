package binarysearch

import "cmp"

type inorderIterator[T cmp.Ordered] struct {
	tree     *Tree[T]
	node     *TreeNode[T]
	finished bool
}

func NewInOrderIterator[T cmp.Ordered](tree *Tree[T]) TreeIterator[T] {
	return &inorderIterator[T]{
		tree: tree,
	}
}

func (i *inorderIterator[T]) HasNext() bool {
	if i.finished {
		return false
	}

	if i.node == nil {
		i.node = i.tree.root
		for i.node != nil {
			if i.node.Left == nil {
				return true
			}

			i.node = i.node.Left
		}
	}

	if i.node.Right != nil {
		i.node = i.node.Right
		for i.node.Left != nil {
			i.node = i.node.Left
		}

		return true
	}

	for i.node.Parent != nil {
		node := i.node
		i.node = i.node.Parent

		if node == i.node.Left {
			return true
		}
	}

	i.finished = true
	return false
}

func (i *inorderIterator[T]) Value() T {
	return i.node.Value
}

func (i *inorderIterator[T]) Reset() {
	i.node = nil
	i.finished = false
}
