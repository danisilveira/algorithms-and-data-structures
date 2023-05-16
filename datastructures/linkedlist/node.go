package linkedlist

type Node[T comparable] struct {
	Next  *Node[T]
	Prev  *Node[T]
	Value T

	list LinkedList[T]
}

func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}
