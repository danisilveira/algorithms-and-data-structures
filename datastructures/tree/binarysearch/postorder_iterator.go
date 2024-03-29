package binarysearch

import (
	"cmp"
	"fmt"
	"strings"
)

type postOrderIterator[T cmp.Ordered] struct {
	tree     *Tree[T]
	node     *TreeNode[T]
	finished bool
}

func NewPostOrderIterator[T cmp.Ordered](tree *Tree[T]) TreeIterator[T] {
	return &postOrderIterator[T]{
		tree: tree,
	}
}

func (i *postOrderIterator[T]) HasNext() bool {
	if i.finished {
		return false
	}

	if i.node == nil {
		i.node = i.tree.LeftmostLeaf(i.tree.root)
		return i.node != nil
	}

	if i.node.Parent != nil {
		node := i.node
		i.node = i.node.Parent

		if i.node.Right != nil && node != i.node.Right {
			i.node = i.tree.LeftmostLeaf(i.node.Right)
			return true
		}

		return true
	}

	i.finished = true
	return false
}

func (i *postOrderIterator[T]) Value() T {
	return i.node.Value
}

func (i *postOrderIterator[T]) String() string {
	var builder strings.Builder

	i.Reset()

	for i.HasNext() {
		value := i.Value()
		builder.WriteString(fmt.Sprintf("%v ", value))
	}

	return strings.TrimSpace(builder.String())
}

func (i *postOrderIterator[T]) Reset() {
	i.node = nil
	i.finished = false
}
