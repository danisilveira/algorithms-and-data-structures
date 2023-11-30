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

func (t *Tree[T]) Get(value T) (*TreeNode[T], bool) {
	currentNode := t.root

	for currentNode != nil {
		if value == currentNode.Value {
			break
		}

		if value > currentNode.Value {
			currentNode = currentNode.Right
			continue
		}

		currentNode = currentNode.Left
	}

	return currentNode, currentNode != nil
}

func (t *Tree[T]) Insert(value T) {
	node := NewBinaryNode(value)
	t.InsertNode(node)
}

func (t *Tree[T]) InsertNode(node *TreeNode[T]) {
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

func (t *Tree[T]) Delete(value T) {
	node, ok := t.Get(value)
	if !ok {
		return
	}

	if node.Left == nil {
		if node.Parent == nil {
			t.root = node.Right
			node.Parent = nil

			return
		}

		if node.Parent.Left == node {
			node.Parent.Left = node.Right
		} else {
			node.Parent.Right = node.Right
		}

		if node.Right != nil {
			node.Right.Parent = node.Parent
		}

		return
	}

	if node.Right == nil {
		if node.Parent == nil {
			t.root = node.Left
			node.Parent = nil

			return
		}

		if node.Parent.Left == node {
			node.Parent.Left = node.Left
		} else {
			node.Parent.Right = node.Left
		}

		node.Left.Parent = node.Parent

		return
	}

	nextInOrder := t.LeftmostLeaf(node.Right)
	if nextInOrder.Parent != node {
		if nextInOrder.Parent.Left == nextInOrder {
			nextInOrder.Parent.Left = nextInOrder.Right
		} else {
			nextInOrder.Parent.Right = nextInOrder.Right
		}

		if nextInOrder.Right != nil {
			nextInOrder.Right.Parent = nextInOrder.Parent
		}

		node.Right.Parent = nextInOrder
		nextInOrder.Right = node.Right
	}

	if node.Parent == nil {
		t.root = nextInOrder
	} else if node.Parent.Left == node {
		node.Parent.Left = nextInOrder
	} else {
		node.Parent.Right = nextInOrder
	}

	nextInOrder.Parent = node.Parent
	nextInOrder.Left = node.Left
	nextInOrder.Left.Parent = nextInOrder
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
