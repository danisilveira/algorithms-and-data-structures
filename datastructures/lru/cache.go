package lru

import (
	"github.com/danisilveira/algorithms-and-data-structures/datastructures/hashtable"
	"github.com/danisilveira/algorithms-and-data-structures/datastructures/linkedlist"
)

type Cache[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V) bool
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

type cache[K comparable, V any] struct {
	capacity     uint64
	storage      *hashtable.HashTable[K, *linkedlist.Node[entry[K, V]]]
	evictionList linkedlist.LinkedList[entry[K, V]]
}

func NewCache[K comparable, V any](capacity uint64) Cache[K, V] {
	return &cache[K, V]{
		capacity:     capacity,
		storage:      hashtable.New[K, *linkedlist.Node[entry[K, V]]](capacity),
		evictionList: linkedlist.New[entry[K, V]](),
	}
}

func (c *cache[K, V]) Get(key K) (V, bool) {
	linkedListNode, ok := c.storage.Get(key)
	if !ok {
		var defaultValue V
		return defaultValue, false
	}

	_ = c.evictionList.MoveToFront(linkedListNode)

	return linkedListNode.Value.value, true
}

func (c *cache[K, V]) Set(key K, value V) bool {
	if linkedListNode, ok := c.storage.Get(key); ok {
		linkedListNode.Value.value = value
		_ = c.evictionList.MoveToFront(linkedListNode)
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
	_ = c.evictionList.AddNodeFirst(linkedListNode)
	c.storage.Set(key, linkedListNode)

	return evict
}

func (c *cache[K, V]) shouldEvict() bool {
	return c.evictionList.Len() >= int(c.capacity)
}
