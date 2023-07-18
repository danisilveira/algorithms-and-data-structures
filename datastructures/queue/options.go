package queue

var basePercentage float64 = 100

type Option[T any] func(queue *queue[T])

func WithCapacity[T any](capacity int) Option[T] {
	return func(queue *queue[T]) {
		queue.capacity = capacity
	}
}

func WithGrowFactor[T any](growFactor float64) Option[T] {
	return func(queue *queue[T]) {
		queue.growFactor = int(growFactor * basePercentage)
	}
}

func WithMinimumGrow[T any](minimumGrow int) Option[T] {
	return func(queue *queue[T]) {
		queue.minimumGrow = minimumGrow
	}
}
