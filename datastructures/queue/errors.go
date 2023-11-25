package queue

import "errors"

var (
	ErrQueueIsEmpty           = errors.New("queue is empty")
	ErrQueueCapacityNegative  = errors.New("queue capacity is negative")
	ErrQueueInvalidGrowFactor = errors.New("invalid grow factor")
)
