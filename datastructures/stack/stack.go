package stack

type Stack[T any] struct {
	items    []T
	length   int
	capacity int

	minimumGrow int
	growFactor  int
}

func New[T any](options ...Option[T]) (*Stack[T], error) {
	stack := &Stack[T]{
		capacity:    10,
		minimumGrow: 4,
		growFactor:  2,
	}

	for _, option := range options {
		option(stack)
	}

	if stack.capacity < 0 {
		return nil, ErrStackCapacityNegative
	}

	if stack.growFactor < 2 || stack.growFactor > 100 {
		return nil, ErrStackInvalidGrowFactor
	}

	stack.items = make([]T, stack.capacity)

	return stack, nil
}

func MustNew[T any](options ...Option[T]) *Stack[T] {
	stack, err := New(options...)
	if err != nil {
		panic(err)
	}

	return stack
}

func (s *Stack[T]) Push(value T) error {
	if s.Full() {
		newCapacity := s.capacity * s.growFactor
		if newCapacity < (s.capacity + s.minimumGrow) {
			newCapacity = s.capacity + s.minimumGrow
		}

		s.capacity = newCapacity
		newItems := make([]T, s.capacity)
		copy(newItems, s.items)
		s.items = newItems
	}

	s.items[s.length] = value
	s.length++

	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	var defaultValue T

	if s.Empty() {
		return defaultValue, ErrStackIsEmpty
	}

	s.length--
	item := s.items[s.length]
	s.items[s.length] = defaultValue

	return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Empty() {
		var defaulValue T
		return defaulValue, ErrStackIsEmpty
	}

	return s.items[s.length-1], nil
}

func (s *Stack[T]) Len() int {
	return s.length
}

func (s *Stack[T]) Cap() int {
	return s.capacity
}

func (s *Stack[T]) Full() bool {
	return s.length == s.capacity
}

func (s *Stack[T]) Empty() bool {
	return s.length == 0
}
