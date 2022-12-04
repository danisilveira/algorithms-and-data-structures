package hashtable

type Option[K Key, V any] func(hashTable *HashTable[K, V])

func WithDefaultHashGeneratorFunc[K Key, V any]() Option[K, V] {
	return func(hashTable *HashTable[K, V]) {
		hashTable.hashGenerator = defaultHashGeneratorFunc[K]
	}
}

func WithCustomHashGeneratorFunc[K Key, V any](hashGeneratorFunc HashGeneratorFunc[K]) Option[K, V] {
	return func(hashTable *HashTable[K, V]) {
		hashTable.hashGenerator = hashGeneratorFunc
	}
}
