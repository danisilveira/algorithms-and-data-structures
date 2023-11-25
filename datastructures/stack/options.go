package stack

type Option[T any] func(stack *Stack[T])

func WithCapacity[T any](capacity int) Option[T] {
	return func(stack *Stack[T]) {
		stack.capacity = capacity
	}
}

func WithGrowFactor[T any](growFactor int) Option[T] {
	return func(stack *Stack[T]) {
		stack.growFactor = growFactor
	}
}

func WithMinimumGrow[T any](minimumGrow int) Option[T] {
	return func(stack *Stack[T]) {
		stack.minimumGrow = minimumGrow
	}
}
