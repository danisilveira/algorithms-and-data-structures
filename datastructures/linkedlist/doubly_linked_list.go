package linkedlist

import (
	"fmt"
	"strings"
)

type Doubly[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewDoubly[T any]() *Doubly[T] {
	return &Doubly[T]{}
}

func (ll *Doubly[T]) AddValueFirst(value T) *Node[T] {
	newNode := NewNode(value)
	ll.AddNodeFirst(newNode)

	return newNode
}

func (ll *Doubly[T]) AddNodeFirst(newNode *Node[T]) {
	if ll.Empty() {
		ll.insertNodeToEmptyList(newNode)
		return
	}

	ll.insertNodeBefore(ll.head, newNode)
	ll.head = newNode
}

func (ll *Doubly[T]) AddValueLast(value T) *Node[T] {
	newNode := NewNode(value)
	ll.AddNodeLast(newNode)

	return newNode
}

func (ll *Doubly[T]) AddNodeLast(newNode *Node[T]) {
	if ll.Empty() {
		ll.insertNodeToEmptyList(newNode)
		return
	}

	ll.insertNodeAfter(ll.tail, newNode)
	ll.tail = newNode
}

func (ll *Doubly[T]) AddValueAfter(node *Node[T], value T) *Node[T] {
	newNode := NewNode(value)
	ll.AddNodeAfter(node, newNode)

	return newNode
}

func (ll *Doubly[T]) AddNodeAfter(node, newNode *Node[T]) {
	ll.insertNodeAfter(node, newNode)

	if ll.tail == node {
		ll.tail = newNode
	}
}

func (ll *Doubly[T]) AddValueBefore(node *Node[T], value T) *Node[T] {
	newNode := NewNode(value)
	ll.AddNodeBefore(node, newNode)

	return newNode
}

func (ll *Doubly[T]) AddNodeBefore(node, newNode *Node[T]) {
	ll.insertNodeBefore(node, newNode)

	if ll.head == node {
		ll.head = newNode
	}
}

func (ll *Doubly[T]) MoveToFront(node *Node[T]) {
	if node == ll.head {
		return
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
}

func (ll *Doubly[T]) MoveToBack(node *Node[T]) {
	if node == ll.tail {
		return
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
}

func (ll *Doubly[T]) RemoveNode(node *Node[T]) {
	ll.removeNode(node)
}

func (ll *Doubly[T]) RemoveFirst() error {
	if ll.Empty() {
		return ErrLinkedListIsEmpty
	}

	ll.removeNode(ll.head)

	return nil
}

func (ll *Doubly[T]) RemoveLast() error {
	if ll.Empty() {
		return ErrLinkedListIsEmpty
	}

	ll.removeNode(ll.tail)

	return nil
}

func (ll *Doubly[T]) FrontNode() *Node[T] {
	return ll.head
}

func (ll *Doubly[T]) BackNode() *Node[T] {
	return ll.tail
}

func (ll *Doubly[T]) Len() int {
	return ll.length
}

func (ll *Doubly[T]) Empty() bool {
	return ll.length == 0
}

func (ll Doubly[T]) String() string {
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

func (ll *Doubly[T]) insertNodeToEmptyList(newNode *Node[T]) {
	ll.head = newNode
	ll.tail = newNode
	ll.length++
}

func (ll *Doubly[T]) insertNodeBefore(node, newNode *Node[T]) {
	newNode.Next = node
	newNode.Prev = node.Prev

	if node.Prev != nil {
		node.Prev.Next = newNode
	}

	node.Prev = newNode

	ll.length++
}

func (ll *Doubly[T]) insertNodeAfter(node, newNode *Node[T]) {
	newNode.Prev = node
	newNode.Next = node.Next

	if node.Next != nil {
		node.Next.Prev = newNode
	}

	node.Next = newNode
	ll.length++
}

func (ll *Doubly[T]) removeNode(node *Node[T]) {
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
