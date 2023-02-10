package queue

import "sync"

type synchronized[T any] struct {
	q  Queue[T]
	mu *sync.RWMutex
}

func newSynchronized[T any](q Queue[T]) Queue[T] {
	return &synchronized[T]{
		q:  q,
		mu: &sync.RWMutex{},
	}
}

func (s *synchronized[T]) Enqueue(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.q.Enqueue(value)
}

func (s *synchronized[T]) Dequeue() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.q.Dequeue()
}

func (s *synchronized[T]) Peek() (T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.q.Peek()
}

func (s *synchronized[T]) Length() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.q.Length()
}

func (s *synchronized[T]) Capacity() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.q.Capacity()
}

func (s *synchronized[T]) Empty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.q.Empty()
}

func (s *synchronized[T]) Full() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.q.Full()
}

func (s *synchronized[T]) Synchronized() Queue[T] {
	return s
}

func (s *synchronized[T]) IsSynchronized() bool {
	return true
}
