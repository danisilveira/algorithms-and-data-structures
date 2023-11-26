package binarysearch

import "cmp"

type TreeIterator[T cmp.Ordered] interface {
	HasNext() bool
	Value() T
	Reset()
}
