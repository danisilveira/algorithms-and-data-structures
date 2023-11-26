package queue

type Circular[T any] struct {
	items    []T
	head     int
	tail     int
	length   int
	capacity int
}

func NewCircular[T any](options ...CircularOption[T]) (*Circular[T], error) {
	queue := &Circular[T]{
		capacity: 32,
	}

	for _, option := range options {
		option(queue)
	}

	queue.items = make([]T, queue.capacity)

	return queue, nil
}

func MustNewCircular[T any](options ...CircularOption[T]) *Circular[T] {
	queue, err := NewCircular(options...)
	if err != nil {
		panic(err)
	}

	return queue
}

func (q *Circular[T]) Enqueue(value T) {
	if q.Full() {
		_, _ = q.Dequeue()
	}

	q.items[q.tail] = value
	q.tail = q.next(q.tail)

	q.length++
}

func (q *Circular[T]) Dequeue() (T, error) {
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

func (q *Circular[T]) Peek() (T, error) {
	if q.Empty() {
		var defaultValue T
		return defaultValue, ErrQueueIsEmpty
	}

	return q.items[q.head], nil
}

func (q *Circular[T]) Len() int {
	return q.length
}

func (q *Circular[T]) Cap() int {
	return q.capacity
}

func (q *Circular[T]) Full() bool {
	return q.length == q.capacity
}

func (q *Circular[T]) Empty() bool {
	return q.length == 0
}

func (q *Circular[T]) next(index int) int {
	return (index + 1) % q.capacity
}
