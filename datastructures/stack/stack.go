package stack

import "errors"

var (
	ErrStackOverflow = errors.New("stack overflow")
	ErrStackIsEmpty  = errors.New("stack is empty")
)

type stack[T any] struct {
	items    []T
	top      int64
	capacity int64
}

func New[T any](capacity int64) *stack[T] {
	return &stack[T]{
		items:    make([]T, capacity),
		top:      -1,
		capacity: capacity,
	}
}

func (s *stack[T]) Push(value T) error {
	if (s.top + 1) == s.capacity {
		return ErrStackOverflow
	}

	s.top++
	s.items[s.top] = value

	return nil
}

func (s *stack[T]) Pop() (T, error) {
	if s.top <= -1 {
		var defaultValue T
		return defaultValue, ErrStackIsEmpty
	}

	item := s.items[s.top]
	s.top--

	return item, nil
}

func (s *stack[T]) Empty() bool {
	return s.top == -1
}

func (s *stack[T]) Full() bool {
	return (s.top + 1) == s.capacity
}
