package queue_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("it should enqueue and dequeue items correctly", func(t *testing.T) {
		q := queue.MustNewCircular(queue.CircularWithCapacity[int](3))
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
	})

	t.Run("it should be circular", func(t *testing.T) {
		q := queue.MustNewCircular(
			queue.CircularWithCapacity[int](3),
		)

		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		assert.Equal(t, 3, q.Len())
		assert.Equal(t, 3, q.Cap())
		assert.True(t, q.Full())

		q.Enqueue(4)
		q.Enqueue(5)
		q.Enqueue(6)

		assert.Equal(t, 3, q.Len())
		assert.Equal(t, 3, q.Cap())
		assert.True(t, q.Full())

		four, err := q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 4, four)

		five, err := q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 5, five)

		six, err := q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 6, six)
	})

	t.Run("it should return an error when an item is dequeued from an empty queue", func(t *testing.T) {
		q := queue.MustNewCircular[string]()
		assert.True(t, q.Empty())

		_, err := q.Dequeue()
		assert.Error(t, err)
		assert.ErrorIs(t, err, queue.ErrQueueIsEmpty)
	})
}
