package queue

type Circular[T any] struct {
	items    []T
	head     int
	tail     int
	length   int
	capacity int

	minimumGrow int
	growFactor  int
}

func NewCircular[T any](options ...CircularOption[T]) (*Circular[T], error) {
	queue := &Circular[T]{
		capacity:    32,
		growFactor:  2,
		minimumGrow: 4,
	}

	for _, option := range options {
		option(queue)
	}

	queue.items = make([]T, queue.capacity)

	if queue.capacity < 0 {
		return nil, ErrQueueCapacityNegative
	}

	if queue.growFactor < 2 || queue.growFactor > 100 {
		return nil, ErrQueueInvalidGrowFactor
	}

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
		newCapacity := q.capacity * q.growFactor
		if newCapacity < (q.capacity + q.minimumGrow) {
			newCapacity = q.capacity + q.minimumGrow
		}

		q.setCapacity(newCapacity)
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

func (q *Circular[T]) setCapacity(capacity int) {
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

func (q *Circular[T]) next(index int) int {
	return (index + 1) % q.capacity
}
