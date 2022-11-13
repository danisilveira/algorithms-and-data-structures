package queue

import "errors"

var (
	ErrQueueIsFull  = errors.New("queue is full")
	ErrQueueIsEmpty = errors.New("queue is empty")
)

type queue[T any] struct {
	items    []T
	front    int
	back     int
	capacity int
	count    int
}

func New[T any](capacity int) *queue[T] {
	return &queue[T]{
		items:    make([]T, capacity),
		front:    0,
		back:     0,
		capacity: capacity,
	}
}

func (q *queue[T]) Enqueue(value T) error {
	if q.Full() {
		return ErrQueueIsFull
	}

	q.items[q.back] = value
	q.back = q.next(q.back)

	q.count++

	return nil
}

func (q *queue[T]) Dequeue() (T, error) {
	if q.Empty() {
		var defaultValue T
		return defaultValue, ErrQueueIsEmpty
	}

	item := q.items[q.front]
	q.front = q.next(q.front)
	q.count--

	return item, nil
}

func (q *queue[T]) Capacity() int {
	return q.capacity
}

func (q *queue[T]) Length() int {
	return q.count
}

func (q *queue[T]) Empty() bool {
	return q.count == 0
}

func (q *queue[T]) Full() bool {
	return q.count == q.capacity
}

func (q *queue[T]) next(index int) int {
	return (index + 1) % q.capacity
}
