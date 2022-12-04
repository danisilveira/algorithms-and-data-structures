package hashtable

import (
	"hash/fnv"
	"sync"
)

type Key interface {
	~string
}

type HashGeneratorFunc[K Key] func(key K) int64

type item[K Key, V any] struct {
	key   K
	value V
	next  *item[K, V]
}

type HashTable[K Key, V any] struct {
	items         []*item[K, V]
	capacity      int64
	hashGenerator HashGeneratorFunc[K]
	mu            *sync.RWMutex
}

func New[K Key, V any](capacity int64, options ...Option[K, V]) *HashTable[K, V] {
	hashTable := HashTable[K, V]{
		items:         make([]*item[K, V], capacity, capacity),
		capacity:      capacity,
		hashGenerator: defaultHashGeneratorFunc[K],
		mu:            &sync.RWMutex{},
	}

	for _, option := range options {
		option(&hashTable)
	}

	return &hashTable
}

func (h *HashTable[K, V]) Get(key K) (V, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	hash := h.hashGenerator(key) % h.capacity

	item := h.items[hash]
	if item != nil && item.key == key {
		return item.value, true
	}

	for item != nil && item.next != nil {
		item = item.next

		if item.key == key {
			return item.value, true
		}
	}

	var defaultValue V
	return defaultValue, false
}

func (h *HashTable[K, V]) Set(key K, value V) {
	h.mu.Lock()
	defer h.mu.Unlock()

	hash := h.hashGenerator(key) % h.capacity

	newItem := &item[K, V]{
		key:   key,
		value: value,
	}

	element := h.items[hash]
	if element == nil {
		h.items[hash] = newItem
		return
	}

	if element.key == key {
		element.value = value
		return
	}

	for element.next != nil {
		element = element.next

		if element.key == key {
			element.value = value
			return
		}
	}

	element.next = newItem
}

func defaultHashGeneratorFunc[K Key](key K) int64 {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int64(hash.Sum32())
}
