package hashtable

type Option[K comparable, V any] func(hashTable *HashTable[K, V])

func WithDefaultHashGeneratorFunc[K comparable, V any]() Option[K, V] {
	return func(hashTable *HashTable[K, V]) {
		hashTable.hashGenerator = defaultHashGeneratorFunc[K]
	}
}

func WithCustomHashGeneratorFunc[K comparable, V any](hashGeneratorFunc HashGeneratorFunc[K]) Option[K, V] {
	return func(hashTable *HashTable[K, V]) {
		hashTable.hashGenerator = hashGeneratorFunc
	}
}
