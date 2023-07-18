package stack

var basePercentage float64 = 100

type Option[T any] func(stack *stack[T])

func WithCapacity[T any](capacity int) Option[T] {
	return func(stack *stack[T]) {
		stack.capacity = capacity
	}
}

func WithGrowFactor[T any](growFactor float64) Option[T] {
	return func(stack *stack[T]) {
		stack.growFactor = int(growFactor * basePercentage)
	}
}

func WithMinimumGrow[T any](minimumGrow int) Option[T] {
	return func(stack *stack[T]) {
		stack.minimumGrow = minimumGrow
	}
}
