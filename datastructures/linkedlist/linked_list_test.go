package linkedlist_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	list := linkedlist.New[int]()
	assert.True(t, list.Empty())

	list.AddValueLast(1)
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, "[1]", list.String())

	three := linkedlist.NewNode(3)
	list.AddNodeLast(three)
	assert.Equal(t, 2, list.Len())
	assert.Equal(t, "[1 3]", list.String())

	list.AddValueBefore(three, 2)
	assert.Equal(t, 3, list.Len())
	assert.Equal(t, "[1 2 3]", list.String())

	list.AddValueAfter(three, 4)
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, "[1 2 3 4]", list.String())

	one, err := list.Find(1)
	assert.NoError(t, err)

	zero := linkedlist.NewNode(0)
	list.AddNodeBefore(one, zero)
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, "[0 1 2 3 4]", list.String())

	four, err := list.Find(4)
	assert.NoError(t, err)

	five := linkedlist.NewNode(5)
	list.AddNodeAfter(four, five)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[0 1 2 3 4 5]", list.String())

	list.AddValueFirst(-1)
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 5]", list.String())

	list.AddValueLast(6)
	assert.Equal(t, 8, list.Len())
	assert.Equal(t, "[-1 0 1 2 3 4 5 6]", list.String())

	twoNegatives := linkedlist.NewNode(-2)
	list.AddNodeFirst(twoNegatives)
	assert.Equal(t, 9, list.Len())
	assert.Equal(t, "[-2 -1 0 1 2 3 4 5 6]", list.String())

	seven := linkedlist.NewNode(7)
	list.AddNodeLast(seven)
	assert.Equal(t, 10, list.Len())
	assert.Equal(t, "[-2 -1 0 1 2 3 4 5 6 7]", list.String())

	list.RemoveValue(-1)
	assert.Equal(t, 9, list.Len())
	assert.Equal(t, "[-2 0 1 2 3 4 5 6 7]", list.String())

	list.RemoveFirst()
	assert.Equal(t, 8, list.Len())
	assert.Equal(t, "[0 1 2 3 4 5 6 7]", list.String())

	list.RemoveLast()
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "[0 1 2 3 4 5 6]", list.String())

	list.RemoveNode(five)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[0 1 2 3 4 6]", list.String())

	six, err := list.FindLast(6)
	assert.NoError(t, err)

	list.MoveToFront(six)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[6 0 1 2 3 4]", list.String())

	list.MoveToBack(six)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[0 1 2 3 4 6]", list.String())

	list.MoveToBack(six)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[0 1 2 3 4 6]", list.String())

	head := list.FrontNode()
	list.MoveToBack(head)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[1 2 3 4 6 0]", list.String())

	list.MoveToFront(head)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[0 1 2 3 4 6]", list.String())

	list.MoveToFront(head)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[0 1 2 3 4 6]", list.String())

	three, err = list.FindLast(3)
	assert.NoError(t, err)
	list.MoveToBack(three)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[0 1 2 4 6 3]", list.String())

	two, err := list.Find(2)
	assert.NoError(t, err)
	list.MoveToFront(two)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "[2 0 1 4 6 3]", list.String())

	_, err = list.Find(1000)
	assert.Equal(t, linkedlist.ErrNodeNotFound, err)

	_, err = list.FindLast(1000)
	assert.Equal(t, linkedlist.ErrNodeNotFound, err)

	err = list.RemoveValue(1000)
	assert.Equal(t, linkedlist.ErrNodeNotFound, err)

	tail := list.BackNode()
	assert.Equal(t, 3, tail.Value)
}
