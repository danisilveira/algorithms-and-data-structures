package linkedlist

import "errors"

var (
	ErrNodeNotFound      = errors.New("linked list node not found")
	ErrLinkedListIsEmpty = errors.New("linked list is empty")
)
