package stack

import "errors"

var (
	ErrStackOverflow = errors.New("stack overflow")
	ErrStackIsEmpty  = errors.New("stack is empty")
)

type Stack[T any] interface {
	Push(T) error
	Pop() (T, error)
	Top() T
	Length() int
	Capacity() int
	Full() bool
	Empty() bool
}

type stack[T any] struct {
	items    []T
	top      int
	capacity int
}

func New[T any](capacity int) Stack[T] {
	return &stack[T]{
		items:    make([]T, capacity),
		top:      0,
		capacity: capacity,
	}
}

func (s *stack[T]) Push(value T) error {
	if s.Full() {
		return ErrStackOverflow
	}

	s.items[s.top] = value
	s.top++

	return nil
}

func (s *stack[T]) Pop() (T, error) {
	if s.Empty() {
		var defaultValue T
		return defaultValue, ErrStackIsEmpty
	}

	s.top--
	item := s.items[s.top]

	return item, nil
}

func (s *stack[T]) Top() T {
	return s.items[s.top-1]
}

func (s *stack[T]) Length() int {
	return s.top
}

func (s *stack[T]) Capacity() int {
	return s.capacity
}

func (s *stack[T]) Full() bool {
	return s.top == s.capacity
}

func (s *stack[T]) Empty() bool {
	return s.top == 0
}
