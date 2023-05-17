package lru

import (
	"github.com/danisilveira/algorithms-and-data-structures/datastructures/hashtable"
	"github.com/danisilveira/algorithms-and-data-structures/datastructures/linkedlist"
)

type Cache[K hashtable.Key, V comparable] interface {
	Get(key K) (V, bool)
	Set(key K, value V) bool
}

type entry[K hashtable.Key, V comparable] struct {
	key   K
	value V
}

type cache[K hashtable.Key, V comparable] struct {
	capacity     int64
	storage      *hashtable.HashTable[K, *linkedlist.Node[entry[K, V]]]
	evictionList linkedlist.LinkedList[entry[K, V]]
}

func NewCache[K hashtable.Key, V comparable](capacity int64) Cache[K, V] {
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

	c.evictionList.MoveToFront(linkedListNode)

	return linkedListNode.Value.value, true
}

func (c *cache[K, V]) Set(key K, value V) bool {
	if linkedListNode, ok := c.storage.Get(key); ok {
		linkedListNode.Value.value = value
		c.evictionList.MoveToFront(linkedListNode)
		return false
	}

	evict := c.shouldEvict()
	if evict {
		linkedListNode := c.evictionList.BackNode()
		c.evictionList.RemoveLast()
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

func (c *cache[K, V]) shouldEvict() bool {
	return c.evictionList.Len() >= int(c.capacity)
}
