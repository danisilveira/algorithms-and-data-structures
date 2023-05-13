package queue_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := queue.MustNew(queue.WithCapacity[int](3))
	assert.True(t, q.Empty())
	assert.Equal(t, 0, q.Len())
	assert.Equal(t, 3, q.Cap())

	q.Enqueue(1)
	assert.Equal(t, 1, q.Len())
	assert.False(t, q.Empty())

	one, err := q.Dequeue()
	assert.Equal(t, 1, one)
	assert.Nil(t, err)
	assert.Equal(t, 0, q.Len())

	assert.False(t, q.IsSynchronized())
	q = q.Synchronized()
	assert.True(t, q.IsSynchronized())

	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	assert.True(t, q.Full())
	assert.Equal(t, 3, q.Len())

	two, err := q.Dequeue()
	assert.Equal(t, 2, two)
	assert.Nil(t, err)

	three, err := q.Dequeue()
	assert.Equal(t, 3, three)
	assert.Nil(t, err)

	four, err := q.Dequeue()
	assert.Equal(t, 4, four)
	assert.Nil(t, err)

	assert.True(t, q.Empty())
	assert.Equal(t, 0, q.Len())
}

func TestQueue_ShouldBeAbleToIncreaseItsCapacity(t *testing.T) {
	q := queue.MustNew(
		queue.WithCapacity[rune](1),
	)
	assert.True(t, q.Empty())

	q.Enqueue('a')
	assert.True(t, q.Full())
	q.Enqueue('b')
	assert.False(t, q.Full())

	assert.Equal(t, 2, q.Len())
	assert.Equal(t, 5, q.Cap())

	a, err := q.Peek()
	assert.Equal(t, 'a', a)
	assert.Nil(t, err)

	a, err = q.Dequeue()
	assert.Equal(t, 'a', a)
	assert.Nil(t, err)

	b, err := q.Dequeue()
	assert.Equal(t, 'b', b)
	assert.Nil(t, err)
}

func TestQueue_ShouldReturnAnErrorWhenAnItemIsDequeuedFromAnEmptyQueue(t *testing.T) {
	q := queue.MustNew[string]()
	assert.True(t, q.Empty())

	_, err := q.Dequeue()
	assert.Error(t, err)
	assert.ErrorIs(t, err, queue.ErrQueueIsEmpty)
}
