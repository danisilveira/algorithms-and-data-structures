package binarysearch

import (
	"cmp"
)

type Tree[T cmp.Ordered] struct {
	root *TreeNode[T]
}

func NewTree[T cmp.Ordered]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) AddValue(value T) {
	node := NewBinaryNode(value)
	t.AddNode(node)
}

func (t *Tree[T]) AddNode(node *TreeNode[T]) {
	if t.root == nil {
		t.root = node
		return
	}

	currentNode := t.root
	for currentNode != nil {
		if node.Value >= currentNode.Value {
			if currentNode.Right == nil {
				node.Parent = currentNode
				currentNode.Right = node
				break
			}

			currentNode = currentNode.Right
			continue
		}

		if currentNode.Left == nil {
			node.Parent = currentNode
			currentNode.Left = node
			break
		}

		currentNode = currentNode.Left
	}
}

func (t *Tree[T]) LeftmostLeaf(node *TreeNode[T]) *TreeNode[T] {
	if node == nil {
		return nil
	}

	for node.Left != nil {
		node = node.Left
	}

	return node
}

func (t *Tree[T]) RightmostLeaf(node *TreeNode[T]) *TreeNode[T] {
	if node == nil {
		return nil
	}

	for node.Right != nil {
		node = node.Right
	}

	return node
}

func (t *Tree[T]) InOrder() TreeIterator[T] {
	return NewInOrderIterator(t)
}

func (t *Tree[T]) PreOrder() TreeIterator[T] {
	return NewPreOrderIterator(t)
}

func (t *Tree[T]) PostOrder() TreeIterator[T] {
	return NewPostOrderIterator(t)
}
