package stack_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := stack.MustNew[int]()
	assert.True(t, stack.Empty())
	assert.Equal(t, 0, stack.Len())
	assert.Equal(t, 10, stack.Cap())

	assert.False(t, stack.IsSynchronized())
	stack = stack.Synchronized()
	assert.True(t, stack.IsSynchronized())

	assert.Nil(t, stack.Push(1))

	top, err := stack.Peek()
	assert.Equal(t, 1, top)
	assert.Nil(t, err)
	assert.Equal(t, 1, stack.Len())

	one, err := stack.Pop()
	assert.Equal(t, 1, one)
	assert.Nil(t, err)
	assert.Equal(t, 0, stack.Len())

	assert.Nil(t, stack.Push(2))
	assert.Nil(t, stack.Push(3))

	top, err = stack.Peek()
	assert.Equal(t, 3, top)
	assert.Nil(t, err)

	assert.Equal(t, 2, stack.Len())

	three, err := stack.Pop()
	assert.Equal(t, 3, three)
	assert.Nil(t, err)

	two, err := stack.Pop()
	assert.Equal(t, 2, two)
	assert.Nil(t, err)

	assert.True(t, stack.Empty())
	assert.Equal(t, 0, stack.Len())
}

func TestStack_ShouldBeAbleToIncreaseItsCapacity(t *testing.T) {
	s := stack.MustNew(stack.WithCapacity[int](1))
	assert.Nil(t, s.Push(1))

	assert.True(t, s.Full())

	assert.Nil(t, s.Push(2))

	assert.Equal(t, 2, s.Len())
	assert.Equal(t, 5, s.Cap())
}

func TestStack_ShouldReturnAnErrorWhenAnItemIsPoppedFromAnEmptyStack(t *testing.T) {
	s := stack.MustNew[int]()

	assert.True(t, s.Empty())

	_, err := s.Pop()

	assert.Error(t, err)
	assert.ErrorIs(t, err, stack.ErrStackIsEmpty)
}
