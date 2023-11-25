package stack

import "errors"

var (
	ErrStackIsEmpty           = errors.New("stack is empty")
	ErrStackCapacityNegative  = errors.New("stack capacity is negative")
	ErrStackInvalidGrowFactor = errors.New("invalid grow factor")
)
