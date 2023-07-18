package linkedlist_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/linkedlist"
	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestLinkedList(t *testing.T) {
	list := linkedlist.New[int]()
	assert.True(t, list.Empty())

	one := list.AddValueLast(1)
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, "[1]", list.String())

	three := linkedlist.NewNode(3)
	assert.NoError(t, list.AddNodeLast(three))
	assert.Equal(t, 2, list.Len())
	assert.Equal(t, "[1 3]", list.String())

	two := list.AddValueBefore(three, 2)
	assert.Equal(t, 3, list.Len())
	assert.Equal(t, "[1 2 3]", list.String())

	four := list.AddValueAfter(three, 4)
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, "[1 2 3 4]", list.String())

	zero := linkedlist.NewNode(0)
	assert.NoError(t, list.AddNodeBefore(one, zero))
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, "[0 1 2 3 4]", list.String())

	five := linkedlist.NewNode(5)
	assert.NoError(t, list.AddNodeAfter(four, five))
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[0 1 2 3 4 5]", list.String())

	list.AddValueFirst(-1)
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 5]", list.String())

	six := list.AddValueLast(6)
	assert.Equal(t, 8, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 5 6]", list.String())

	twoNegatives := linkedlist.NewNode(-2)
	assert.NoError(t, list.AddNodeFirst(twoNegatives))
	assert.Equal(t, 9, list.Len())
	assert.Equal(t, "[-2 -1 0 1 2 3 4 5 6]", list.String())

	seven := linkedlist.NewNode(7)
	assert.NoError(t, list.AddNodeLast(seven))
	assert.Equal(t, 10, list.Len())
	assert.Equal(t, "[-2 -1 0 1 2 3 4 5 6 7]", list.String())

	assert.NoError(t, list.RemoveFirst())
	assert.Equal(t, 9, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 5 6 7]", list.String())

	assert.NoError(t, list.RemoveLast())
	assert.Equal(t, 8, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 5 6]", list.String())

	assert.NoError(t, list.RemoveNode(five))
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 6]", list.String())

	assert.NoError(t, list.MoveToFront(six))
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[6 -1 0 1 2 3 4]", list.String())

	assert.NoError(t, list.MoveToBack(six))
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 6]", list.String())

	assert.NoError(t, list.MoveToBack(six))
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 6]", list.String())

	head := list.FrontNode()
	assert.NoError(t, list.MoveToBack(head))
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[0 1 2 3 4 6 -1]", list.String())

	assert.NoError(t, list.MoveToFront(head))
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 6]", list.String())

	assert.NoError(t, list.MoveToFront(head))
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 6]", list.String())

	assert.NoError(t, list.MoveToBack(three))
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[-1 0 1 2 4 6 3]", list.String())

	assert.NoError(t, list.MoveToFront(two))
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[2 -1 0 1 4 6 3]", list.String())

	tail := list.BackNode()
	assert.Equal(t, 3, tail.Value)
}
