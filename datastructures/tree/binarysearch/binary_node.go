package binarysearch

type TreeNode[T comparable] struct {
	Parent *TreeNode[T]
	Left   *TreeNode[T]
	Right  *TreeNode[T]
	Value  T
}

func NewBinaryNode[T comparable](value T) *TreeNode[T] {
	return &TreeNode[T]{
		Value: value,
	}
}
