package hashtable

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type HashGeneratorFunc[K comparable] func(key K) uint64

type item[K comparable, V any] struct {
	key   K
	value V
	next  *item[K, V]
}

type HashTable[K comparable, V any] struct {
	items         []*item[K, V]
	capacity      uint64
	hashGenerator HashGeneratorFunc[K]
	mu            *sync.RWMutex
}

func New[K comparable, V any](capacity uint64, options ...Option[K, V]) *HashTable[K, V] {
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

func (h *HashTable[K, V]) Delete(key K) {
	h.mu.Lock()
	defer h.mu.Unlock()

	hash := h.hashGenerator(key) % h.capacity

	element := h.items[hash]
	if element == nil {
		return
	}

	if element.next == nil {
		h.items[hash] = &item[K, V]{}
		return
	}

	if key == element.key {
		h.items[hash] = element.next
		return
	}

	previousElement := element
	currentElement := element.next
	for currentElement != nil {
		if key == currentElement.key {
			break
		}

		previousElement = currentElement
		currentElement = currentElement.next
	}

	if currentElement == nil {
		return
	}

	previousElement.next = currentElement.next
}

func defaultHashGeneratorFunc[K comparable](key K) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	return h.Sum64()
}
