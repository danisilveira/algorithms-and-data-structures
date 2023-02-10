package stack

import "errors"

var (
	ErrStackIsEmpty           = errors.New("stack is empty")
	ErrStackCapacityNegative  = errors.New("stack capacity is negative")
	ErrStackInvalidGrowFactor = errors.New("invalid grow factor")
)

type Stack[T any] interface {
	Push(T) error
	Pop() (T, error)
	Peek() (T, error)
	Length() int
	Capacity() int
	Full() bool
	Empty() bool
	Synchronized() Stack[T]
	IsSynchronized() bool
}

type stack[T any] struct {
	items    []T
	length   int
	capacity int

	minimumGrow int
	growFactor  int
}

func New[T any](options ...Option[T]) (Stack[T], error) {
	stack := &stack[T]{
		capacity:    10,
		minimumGrow: 4,
		growFactor:  200, // 2.0
	}

	for _, option := range options {
		option(stack)
	}

	if stack.capacity < 0 {
		return nil, ErrStackCapacityNegative
	}

	if stack.growFactor < 100 || stack.growFactor > 1000 {
		return nil, ErrStackInvalidGrowFactor
	}

	stack.items = make([]T, stack.capacity)

	return stack, nil
}

func MustNew[T any](options ...Option[T]) Stack[T] {
	stack, err := New(options...)
	if err != nil {
		panic(err)
	}

	return stack
}

func (s *stack[T]) Push(value T) error {
	if s.Full() {
		newCapacity := s.capacity * (s.growFactor / 100)
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

func (s *stack[T]) Pop() (T, error) {
	var defaultValue T

	if s.Empty() {
		return defaultValue, ErrStackIsEmpty
	}

	s.length--
	item := s.items[s.length]
	s.items[s.length] = defaultValue

	return item, nil
}

func (s *stack[T]) Peek() (T, error) {
	if s.Empty() {
		var defaulValue T
		return defaulValue, ErrStackIsEmpty
	}

	return s.items[s.length-1], nil
}

func (s *stack[T]) Length() int {
	return s.length
}

func (s *stack[T]) Capacity() int {
	return s.capacity
}

func (s *stack[T]) Full() bool {
	return s.length == s.capacity
}

func (s *stack[T]) Empty() bool {
	return s.length == 0
}

func (s *stack[T]) Synchronized() Stack[T] {
	return newSynchronized[T](s)
}

func (s *stack[T]) IsSynchronized() bool {
	return false
}
