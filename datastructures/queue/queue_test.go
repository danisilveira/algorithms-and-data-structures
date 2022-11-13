package queue_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := queue.New[int](3)
	assert.True(t, queue.Empty())
	assert.Equal(t, 0, queue.Length())
	assert.Equal(t, 3, queue.Capacity())

	assert.Nil(t, queue.Enqueue(1))
	assert.Equal(t, 1, queue.Length())
	assert.False(t, queue.Empty())

	one, err := queue.Dequeue()
	assert.Equal(t, 1, one)
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Length())

	assert.Nil(t, queue.Enqueue(2))
	assert.Nil(t, queue.Enqueue(3))
	assert.Nil(t, queue.Enqueue(4))

	assert.True(t, queue.Full())
	assert.Equal(t, 3, queue.Length())

	two, err := queue.Dequeue()
	assert.Equal(t, 2, two)
	assert.Nil(t, err)

	three, err := queue.Dequeue()
	assert.Equal(t, 3, three)
	assert.Nil(t, err)

	four, err := queue.Dequeue()
	assert.Equal(t, 4, four)
	assert.Nil(t, err)

	assert.True(t, queue.Empty())
	assert.Equal(t, 0, queue.Length())
}

func TestQueue_ShouldReturnAnErrorWhenAnItemIsEnqueueIntoAFullQueue(t *testing.T) {
	q := queue.New[rune](1)
	assert.True(t, q.Empty())

	assert.Nil(t, q.Enqueue('a'))

	assert.True(t, q.Full())

	err := q.Enqueue('b')
	assert.Error(t, err)
	assert.ErrorIs(t, err, queue.ErrQueueIsFull)
}

func TestQueue_ShouldReturnAnErrorWhenAnItemIsDequeueFromAnEmptyQueue(t *testing.T) {
	q := queue.New[string](1)
	assert.True(t, q.Empty())

	_, err := q.Dequeue()
	assert.Error(t, err)
	assert.ErrorIs(t, err, queue.ErrQueueIsEmpty)
}
