package cache

import (
	"github.com/danisilveira/algorithms-and-data-structures/datastructures/hashtable"
	"github.com/danisilveira/algorithms-and-data-structures/datastructures/linkedlist"
)

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRU[K comparable, V any] struct {
	capacity     uint64
	storage      *hashtable.HashTable[K, *linkedlist.Node[entry[K, V]]]
	evictionList *linkedlist.Doubly[entry[K, V]]
}

func NewLRU[K comparable, V any](capacity uint64) *LRU[K, V] {
	return &LRU[K, V]{
		capacity:     capacity,
		storage:      hashtable.New[K, *linkedlist.Node[entry[K, V]]](capacity),
		evictionList: linkedlist.NewDoubly[entry[K, V]](),
	}
}

func (c *LRU[K, V]) Get(key K) (V, bool) {
	linkedListNode, ok := c.storage.Get(key)
	if !ok {
		var defaultValue V
		return defaultValue, false
	}

	c.evictionList.MoveToFront(linkedListNode)

	return linkedListNode.Value.value, true
}

func (c *LRU[K, V]) Set(key K, value V) bool {
	if linkedListNode, ok := c.storage.Get(key); ok {
		linkedListNode.Value.value = value
		c.evictionList.MoveToFront(linkedListNode)
		return false
	}

	evict := c.shouldEvict()
	if evict {
		linkedListNode := c.evictionList.BackNode()
		_ = c.evictionList.RemoveLast()
		c.storage.Delete(linkedListNode.Value.key)
	}

	linkedListNode := linkedlist.NewNode(entry[K, V]{
		key:   key,
		value: value,
	})
	c.evictionList.AddNodeFirst(linkedListNode)
	c.storage.Set(key, linkedListNode)

	return evict
}

func (c *LRU[K, V]) shouldEvict() bool {
	return c.evictionList.Len() >= int(c.capacity)
}
