package queue

import (
	"errors"
)

var (
	ErrQueueIsEmpty           = errors.New("queue is empty")
	ErrQueueCapacityNegative  = errors.New("queue capacity is negative")
	ErrQueueInvalidGrowFactor = errors.New("invalid grow factor")
)

type Queue[T any] interface {
	Enqueue(value T)
	Dequeue() (T, error)
	Peek() (T, error)
	Len() int
	Cap() int
	Empty() bool
	Full() bool
	IsSynchronized() bool
	Synchronized() Queue[T]
}

type queue[T any] struct {
	items    []T
	head     int
	tail     int
	length   int
	capacity int

	minimumGrow int
	growFactor  int
}

func New[T any](options ...Option[T]) (Queue[T], error) {
	queue := &queue[T]{
		capacity:    32,
		growFactor:  200, // 2.0
		minimumGrow: 4,
	}

	for _, option := range options {
		option(queue)
	}

	queue.items = make([]T, queue.capacity)

	if queue.capacity < 0 {
		return nil, ErrQueueCapacityNegative
	}

	if queue.growFactor < 100 || queue.growFactor > 1000 {
		return nil, ErrQueueInvalidGrowFactor
	}

	return queue, nil
}

func MustNew[T any](options ...Option[T]) Queue[T] {
	queue, err := New(options...)
	if err != nil {
		panic(err)
	}

	return queue
}

func (q *queue[T]) Enqueue(value T) {
	if q.Full() {
		newCapacity := q.capacity * (q.growFactor / 100)
		if newCapacity < (q.capacity + q.minimumGrow) {
			newCapacity = q.capacity + q.minimumGrow
		}

		q.setCapacity(newCapacity)
	}

	q.items[q.tail] = value
	q.tail = q.next(q.tail)

	q.length++
}

func (q *queue[T]) Dequeue() (T, error) {
	var defaultValue T

	if q.Empty() {
		return defaultValue, ErrQueueIsEmpty
	}

	removed := q.items[q.head]
	q.items[q.head] = defaultValue
	q.head = q.next(q.head)
	q.length--

	return removed, nil
}

func (q *queue[T]) Peek() (T, error) {
	if q.Empty() {
		var defaultValue T
		return defaultValue, ErrQueueIsEmpty
	}

	return q.items[q.head], nil
}

func (q *queue[T]) Len() int {
	return q.length
}

func (q *queue[T]) Cap() int {
	return q.capacity
}

func (q *queue[T]) Empty() bool {
	return q.length == 0
}

func (q *queue[T]) Full() bool {
	return q.length == q.capacity
}

func (q *queue[T]) Synchronized() Queue[T] {
	return newSynchronized[T](q)
}

func (q *queue[T]) IsSynchronized() bool {
	return false
}

func (q *queue[T]) setCapacity(capacity int) {
	newItems := make([]T, capacity)
	if q.length > 0 {
		if q.head < q.tail {
			copy(newItems, q.items)
		} else {
			copy(newItems, q.items[q.head:])
			copy(newItems[q.head+1:], q.items[:q.tail])
		}
	}

	q.items = newItems
	q.capacity = capacity
	q.head = 0
	if q.length == capacity {
		q.tail = 0
	} else {
		q.tail = q.length
	}
}

func (q *queue[T]) next(index int) int {
	return (index + 1) % q.capacity
}
