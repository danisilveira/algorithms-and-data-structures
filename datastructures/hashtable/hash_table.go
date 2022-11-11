package hashtable

import (
	"hash/fnv"
)

type Key interface {
	~string
}

type item[K Key, V any] struct {
	key   K
	value V
	next  *item[K, V]
}

type HashTable[K Key, V any] struct {
	items             []*item[K, V]
	capacity          int64
	hashGeneratorFunc HashGeneratorFunc[K]
}

func New[K Key, V any](capacity int64, options ...Option[K, V]) *HashTable[K, V] {
	hashTable := HashTable[K, V]{
		items:             make([]*item[K, V], capacity, capacity),
		capacity:          capacity,
		hashGeneratorFunc: defaultHashGeneratorFunc[K],
	}

	for _, option := range options {
		option(&hashTable)
	}

	return &hashTable
}

func (hashTable *HashTable[K, V]) Get(key K) (V, bool) {
	hash := hashTable.hashGeneratorFunc(key) % hashTable.capacity

	item := hashTable.items[hash]
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

func (hashTable *HashTable[K, V]) Set(key K, value V) {
	hash := hashTable.hashGeneratorFunc(key) % hashTable.capacity

	newItem := &item[K, V]{
		key:   key,
		value: value,
	}

	element := hashTable.items[hash]
	if element == nil {
		hashTable.items[hash] = newItem
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

type HashGeneratorFunc[K Key] func(key K) int64

func defaultHashGeneratorFunc[K Key](key K) int64 {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int64(hash.Sum32())
}

type Option[K Key, V any] func(hashTable *HashTable[K, V])

func WithDefaultHashGeneratorFunc[K Key, V any]() Option[K, V] {
	return func(hashTable *HashTable[K, V]) {
		hashTable.hashGeneratorFunc = defaultHashGeneratorFunc[K]
	}
}

func WithCustomHashGeneratorFunc[K Key, V any](hashGeneratorFunc HashGeneratorFunc[K]) Option[K, V] {
	return func(hashTable *HashTable[K, V]) {
		hashTable.hashGeneratorFunc = hashGeneratorFunc
	}
}
