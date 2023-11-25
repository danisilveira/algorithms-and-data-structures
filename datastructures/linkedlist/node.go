package linkedlist

type Node[T any] struct {
	Next  *Node[T]
	Prev  *Node[T]
	Value T
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}
