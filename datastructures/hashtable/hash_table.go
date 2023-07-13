package hashtable

import (
	"hash/fnv"
	"sync"
)

type Key interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		string
}

type HashGeneratorFunc[K Key] func(key K) uint64

type item[K Key, V any] struct {
	key   K
	value V
	next  *item[K, V]
}

type HashTable[K Key, V any] struct {
	items         []*item[K, V]
	capacity      uint64
	hashGenerator HashGeneratorFunc[K]
	mu            *sync.RWMutex
}

func New[K Key, V any](capacity uint64, options ...Option[K, V]) *HashTable[K, V] {
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

func defaultHashGeneratorFunc[K Key](key K) uint64 {
	switch k := any(key).(type) {
	case string:
		hash := fnv.New32a()
		hash.Write([]byte(k))
		return uint64(hash.Sum32())
	case int:
		return hash(uint64(k))
	case int8:
		return hash(uint64(k))
	case int16:
		return hash(uint64(k))
	case int32:
		return hash(uint64(k))
	case int64:
		return hash(uint64(k))
	case uint:
		return hash(uint64(k))
	case uint8:
		return hash(uint64(k))
	case uint16:
		return hash(uint64(k))
	case uint32:
		return hash(uint64(k))
	case uint64:
		return hash(uint64(k))
	default:
		panic("something went wrong!")
	}
}

func hash(key uint64) uint64 {
	key ^= key >> 33
	key *= 0xff51afd7ed558ccd
	key ^= key >> 33
	key *= 0xc4ceb9fe1a85ec53
	key ^= key >> 33
	return key
}
