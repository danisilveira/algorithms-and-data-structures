package stack

import "sync"

type synchronized[T any] struct {
	stack Stack[T]
	mu    *sync.RWMutex
}

func newSynchronized[T any](stack Stack[T]) Stack[T] {
	return &synchronized[T]{
		stack: stack,
		mu:    &sync.RWMutex{},
	}
}

func (s *synchronized[T]) Push(value T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.stack.Push(value)
}

func (s *synchronized[T]) Pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.stack.Pop()
}

func (s *synchronized[T]) Peek() (T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.stack.Peek()
}

func (s *synchronized[T]) Length() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.stack.Length()
}

func (s *synchronized[T]) Capacity() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.stack.Capacity()
}

func (s *synchronized[T]) Full() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.stack.Full()
}

func (s *synchronized[T]) Empty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.stack.Empty()
}

func (s *synchronized[T]) Synchronized() Stack[T] {
	return s
}

func (s *synchronized[T]) IsSynchronized() bool {
	return true
}
