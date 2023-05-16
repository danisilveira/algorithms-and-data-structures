package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNodeBelongsToAnotherLinkedList = errors.New("linked list node belongs to another linked list")
	ErrNewNodeIsAlreadyAttached       = errors.New("linked list node is already attached to another linked list")
	ErrNodeNotFound                   = errors.New("linked list node not found")
	ErrLinkedListIsEmpty              = errors.New("linked list is empty")
)

type LinkedList[T comparable] interface {
	AddValueFirst(value T) *Node[T]
	AddNodeFirst(newNode *Node[T]) error

	AddValueLast(value T) *Node[T]
	AddNodeLast(newNode *Node[T]) error

	AddValueAfter(node *Node[T], value T) *Node[T]
	AddNodeAfter(node, newNode *Node[T]) error

	AddValueBefore(node *Node[T], value T) *Node[T]
	AddNodeBefore(node, newNode *Node[T]) error

	MoveToFront(node *Node[T]) error
	MoveToBack(node *Node[T]) error

	Find(value T) (*Node[T], error)
	FindLast(value T) (*Node[T], error)

	RemoveValue(value T) error
	RemoveNode(node *Node[T]) error

	RemoveFirst() error
	RemoveLast() error

	Empty() bool

	FrontNode() *Node[T]
	BackNode() *Node[T]

	Len() int

	String() string
}

type linkedList[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T comparable]() LinkedList[T] {
	return &linkedList[T]{}
}

func (ll *linkedList[T]) AddValueFirst(value T) *Node[T] {
	newNode := NewNode(value)
	ll.AddNodeFirst(newNode)

	return newNode
}

func (ll *linkedList[T]) AddNodeFirst(newNode *Node[T]) error {
	err := ll.validateNewNode(newNode)
	if err != nil {
		return err
	}

	if ll.Empty() {
		ll.insertNodeToEmptyList(newNode)
		return nil
	}

	ll.insertNodeBefore(ll.head, newNode)
	ll.head = newNode

	return nil
}

func (ll *linkedList[T]) AddValueLast(value T) *Node[T] {
	newNode := NewNode(value)
	ll.AddNodeLast(newNode)

	return newNode
}

func (ll *linkedList[T]) AddNodeLast(newNode *Node[T]) error {
	err := ll.validateNewNode(newNode)
	if err != nil {
		return err
	}

	if ll.Empty() {
		ll.insertNodeToEmptyList(newNode)
		return nil
	}

	ll.insertNodeAfter(ll.tail, newNode)
	ll.tail = newNode

	return nil
}

func (ll *linkedList[T]) AddValueAfter(node *Node[T], value T) *Node[T] {
	newNode := NewNode(value)
	ll.AddNodeAfter(node, newNode)

	return newNode
}

func (ll *linkedList[T]) AddNodeAfter(node, newNode *Node[T]) error {
	err := ll.validateNode(node)
	if err != nil {
		return err
	}

	err = ll.validateNewNode(newNode)
	if err != nil {
		return err
	}

	ll.insertNodeAfter(node, newNode)

	if ll.tail == node {
		ll.tail = newNode
	}

	return nil
}

func (ll *linkedList[T]) AddValueBefore(node *Node[T], value T) *Node[T] {
	newNode := NewNode(value)
	ll.AddNodeBefore(node, newNode)

	return newNode
}

func (ll *linkedList[T]) AddNodeBefore(node, newNode *Node[T]) error {
	err := ll.validateNode(node)
	if err != nil {
		return err
	}

	err = ll.validateNewNode(newNode)
	if err != nil {
		return err
	}

	ll.insertNodeBefore(node, newNode)

	if ll.head == node {
		ll.head = newNode
	}

	return nil
}

func (ll *linkedList[T]) MoveToFront(node *Node[T]) error {
	err := ll.validateNode(node)
	if err != nil {
		return err
	}

	if node == ll.head {
		return nil
	}

	if node == ll.tail {
		ll.tail = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Prev = ll.head.Prev
	node.Next = ll.head
	ll.head.Prev = node
	ll.head = node

	return nil
}

func (ll *linkedList[T]) MoveToBack(node *Node[T]) error {
	err := ll.validateNode(node)
	if err != nil {
		return err
	}

	if node == ll.tail {
		return nil
	}

	if node == ll.head {
		ll.head = node.Next
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Prev = ll.tail
	node.Next = ll.tail.Next
	ll.tail.Next = node
	ll.tail = node

	return nil
}

func (ll *linkedList[T]) Find(value T) (*Node[T], error) {
	node := ll.head

	for node != nil {
		if node.Value == value {
			return node, nil
		}

		node = node.Next
	}

	return nil, ErrNodeNotFound
}

func (ll *linkedList[T]) FindLast(value T) (*Node[T], error) {
	node := ll.tail

	for node != nil {
		if node.Value == value {
			return node, nil
		}

		node = node.Prev
	}

	return nil, ErrNodeNotFound
}

func (ll *linkedList[T]) RemoveValue(value T) error {
	node, err := ll.Find(value)
	if err != nil {
		return err
	}

	ll.removeNode(node)

	return nil
}

func (ll *linkedList[T]) RemoveNode(node *Node[T]) error {
	err := ll.validateNode(node)
	if err != nil {
		return err
	}

	ll.removeNode(node)

	return nil
}

func (ll *linkedList[T]) RemoveFirst() error {
	if ll.Empty() {
		return ErrLinkedListIsEmpty
	}

	ll.removeNode(ll.head)

	return nil
}

func (ll *linkedList[T]) RemoveLast() error {
	if ll.Empty() {
		return ErrLinkedListIsEmpty
	}

	ll.removeNode(ll.tail)

	return nil
}

func (ll *linkedList[T]) FrontNode() *Node[T] {
	return ll.head
}

func (ll *linkedList[T]) BackNode() *Node[T] {
	return ll.tail
}

func (ll *linkedList[T]) Empty() bool {
	return ll.length == 0
}

func (ll *linkedList[T]) Len() int {
	return ll.length
}

func (ll linkedList[T]) String() string {
	var builder strings.Builder
	builder.WriteRune('[')

	node := ll.head
	for node != nil {
		builder.WriteString(fmt.Sprintf("%v", node.Value))
		if node.Next != nil {
			builder.WriteRune(' ')
		}

		node = node.Next
	}

	builder.WriteRune(']')
	return builder.String()
}

func (ll *linkedList[T]) validateNode(node *Node[T]) error {
	if node.list != ll {
		return ErrNodeBelongsToAnotherLinkedList
	}

	return nil
}

func (ll *linkedList[T]) validateNewNode(node *Node[T]) error {
	if node.list != nil {
		return ErrNewNodeIsAlreadyAttached
	}

	return nil
}

func (ll *linkedList[T]) insertNodeToEmptyList(newNode *Node[T]) {
	ll.head = newNode
	ll.tail = newNode
	ll.length++

	newNode.list = ll
}

func (ll *linkedList[T]) insertNodeBefore(node, newNode *Node[T]) {
	newNode.Next = node
	newNode.Prev = node.Prev

	if node.Prev != nil {
		node.Prev.Next = newNode
	}

	node.Prev = newNode

	ll.length++

	newNode.list = ll
}

func (ll *linkedList[T]) insertNodeAfter(node, newNode *Node[T]) {
	newNode.Prev = node
	newNode.Next = node.Next

	if node.Next != nil {
		node.Next.Prev = newNode
	}

	node.Next = newNode
	ll.length++

	newNode.list = ll
}

func (ll *linkedList[T]) removeNode(node *Node[T]) {
	if ll.length == 1 && node == ll.head && node == ll.tail {
		ll.head = nil
		ll.tail = nil
		ll.length--

		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node == ll.head {
		ll.head = node.Next
	}

	if node == ll.tail {
		ll.tail = node.Prev
	}

	ll.length--
}
