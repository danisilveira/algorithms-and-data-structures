package queue

type Option[T any] func(queue *queue[T])

func WithCapacity[T any](capacity int) Option[T] {
	return func(queue *queue[T]) {
		queue.capacity = capacity
	}
}

func WithGrowFactor[T any](growFactor float64) Option[T] {
	return func(queue *queue[T]) {
		queue.growFactor = int(growFactor * 100)
	}
}

func WithMinimumGrow[T any](minimumGrow int) Option[T] {
	return func(queue *queue[T]) {
		queue.minimumGrow = minimumGrow
	}
}
