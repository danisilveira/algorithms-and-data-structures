package queue

type CircularOption[T any] func(queue *Circular[T])

func CircularWithCapacity[T any](capacity int) CircularOption[T] {
	return func(queue *Circular[T]) {
		queue.capacity = capacity
	}
}
