package stack_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	st := stack.MustNew[int]()
	assert.True(t, st.Empty())
	assert.Equal(t, 0, st.Len())
	assert.Equal(t, 10, st.Cap())

	assert.False(t, st.IsSynchronized())
	st = st.Synchronized()
	assert.True(t, st.IsSynchronized())

	assert.Nil(t, st.Push(1))

	top, err := st.Peek()
	assert.Equal(t, 1, top)
	assert.Nil(t, err)
	assert.Equal(t, 1, st.Len())

	one, err := st.Pop()
	assert.Equal(t, 1, one)
	assert.Nil(t, err)
	assert.Equal(t, 0, st.Len())

	assert.Nil(t, st.Push(2))
	assert.Nil(t, st.Push(3))

	top, err = st.Peek()
	assert.Equal(t, 3, top)
	assert.Nil(t, err)

	assert.Equal(t, 2, st.Len())

	three, err := st.Pop()
	assert.Equal(t, 3, three)
	assert.Nil(t, err)

	two, err := st.Pop()
	assert.Equal(t, 2, two)
	assert.Nil(t, err)

	assert.True(t, st.Empty())
	assert.Equal(t, 0, st.Len())
}

func TestStack_ShouldBeAbleToIncreaseItsCapacity(t *testing.T) {
	st := stack.MustNew(stack.WithCapacity[int](1))
	assert.Nil(t, st.Push(1))

	assert.True(t, st.Full())

	assert.Nil(t, st.Push(2))

	assert.Equal(t, 2, st.Len())
	assert.Equal(t, 5, st.Cap())
}

func TestStack_ShouldReturnAnErrorWhenAnItemIsPoppedFromAnEmptyStack(t *testing.T) {
	st := stack.MustNew[int]()

	assert.True(t, st.Empty())

	_, err := st.Pop()

	assert.Error(t, err)
	assert.ErrorIs(t, err, stack.ErrStackIsEmpty)
}
