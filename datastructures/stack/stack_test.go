package stack_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := stack.New[int](5)

	assert.Nil(t, stack.Push(1))
	one, err := stack.Pop()
	assert.Equal(t, 1, one)
	assert.Nil(t, err)

	assert.Nil(t, stack.Push(2))
	assert.Nil(t, stack.Push(3))

	three, err := stack.Pop()
	assert.Equal(t, 3, three)
	assert.Nil(t, err)

	two, err := stack.Pop()
	assert.Equal(t, 2, two)
	assert.Nil(t, err)

	assert.True(t, stack.Empty())
}

func TestStack_ShouldReturnAnErrorWhenAnItemIsPushedIntoAFullStack(t *testing.T) {
	s := stack.New[int](1)
	assert.Nil(t, s.Push(1))

	assert.True(t, s.Full())

	err := s.Push(2)
	assert.Error(t, err)
	assert.ErrorIs(t, err, stack.ErrStackOverflow)
}

func TestStack_ShouldReturnAnErrorWhenAnItemIsPoppedFromAnEmptyStack(t *testing.T) {
	s := stack.New[int](1)

	assert.True(t, s.Empty())

	_, err := s.Pop()

	assert.Error(t, err)
	assert.ErrorIs(t, err, stack.ErrStackIsEmpty)
}
