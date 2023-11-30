package binarysearch_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/tree/binarysearch"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	tr := binarysearch.NewTree[int]()
	tr.Insert(48)
	tr.Insert(10)
	tr.Insert(87)
	tr.Insert(74)
	tr.Insert(47)
	tr.Insert(1)
	tr.Insert(31)
	tr.Insert(32)

	assert.Equal(t, "1 10 31 32 47 48 74 87", tr.InOrder().String())

	tr.Delete(10)
	assert.Equal(t, "1 31 32 47 48 74 87", tr.InOrder().String())

	tr.Delete(87)
	assert.Equal(t, "1 31 32 47 48 74", tr.InOrder().String())

	tr.Delete(48)
	assert.Equal(t, "1 31 32 47 74", tr.InOrder().String())
}
