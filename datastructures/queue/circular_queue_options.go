package queue

type CircularOption[T any] func(queue *Circular[T])

func CircularWithCapacity[T any](capacity int) CircularOption[T] {
	return func(queue *Circular[T]) {
		queue.capacity = capacity
	}
}

func CircularWithGrowFactor[T any](growFactor int) CircularOption[T] {
	return func(queue *Circular[T]) {
		queue.growFactor = growFactor
	}
}

func CircularWithMinimumGrow[T any](minimumGrow int) CircularOption[T] {
	return func(queue *Circular[T]) {
		queue.minimumGrow = minimumGrow
	}
}
