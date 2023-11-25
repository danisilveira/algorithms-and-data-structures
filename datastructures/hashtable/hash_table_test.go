package hashtable_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/hashtable"
	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	t.Run("it should be able to put and retrieve some key", func(t *testing.T) {
		hashTable := hashtable.New[string, string](5)

		hashTable.Set("name", "Daniel")
		name, ok := hashTable.Get("name")
		assert.Equal(t, "Daniel", name)
		assert.True(t, ok)
	})

	t.Run("it should be able to deal with collisions correctly", func(t *testing.T) {
		hashGeneratorFunc := func(key string) uint64 {
			return 1
		}

		hashTable := hashtable.New(
			5,
			hashtable.WithCustomHashGeneratorFunc[string, int64](hashGeneratorFunc),
		)

		hashTable.Set("a", 5)
		hashTable.Set("b", 6)

		a, ok := hashTable.Get("a")
		assert.Equal(t, int64(5), a)
		assert.True(t, ok)

		b, ok := hashTable.Get("b")
		assert.Equal(t, int64(6), b)
		assert.True(t, ok)
	})

	t.Run("it should override the value with the same key", func(t *testing.T) {
		hashTable := hashtable.New[string, string](5)

		hashTable.Set("name", "Daniel")
		name, ok := hashTable.Get("name")
		assert.Equal(t, "Daniel", name)
		assert.True(t, ok)

		hashTable.Set("name", "Lucas")
		name, ok = hashTable.Get("name")
		assert.Equal(t, "Lucas", name)
		assert.True(t, ok)
	})

	t.Run("it should override the value with the same key even with a collision", func(t *testing.T) {
		hashGeneratorFunc := func(key string) uint64 {
			return 1
		}

		hashTable := hashtable.New(
			5,
			hashtable.WithCustomHashGeneratorFunc[string, int64](hashGeneratorFunc),
		)

		hashTable.Set("a", 5)
		hashTable.Set("b", 6)

		a, ok := hashTable.Get("a")
		assert.Equal(t, int64(5), a)
		assert.True(t, ok)

		b, ok := hashTable.Get("b")
		assert.Equal(t, int64(6), b)
		assert.True(t, ok)

		hashTable.Set("b", 10)

		b, ok = hashTable.Get("b")
		assert.Equal(t, int64(10), b)
		assert.True(t, ok)
	})

	t.Run("it should return the default value when the key is not found", func(t *testing.T) {
		hashTable := hashtable.New[string, string](5)

		hashTable.Set("name", "Daniel")
		name, ok := hashTable.Get("name")
		assert.Equal(t, "Daniel", name)
		assert.True(t, ok)

		key, ok := hashTable.Get("otherKey")
		assert.Equal(t, "", key)
		assert.False(t, ok)
	})

	t.Run("it should delete a key correctly", func(t *testing.T) {
		hashTable := hashtable.New[string, string](5)

		hashTable.Set("name", "Daniel")
		name, ok := hashTable.Get("name")
		assert.Equal(t, "Daniel", name)
		assert.True(t, ok)

		hashTable.Delete("name")
		name, ok = hashTable.Get("name")
		assert.Equal(t, "", name)
		assert.False(t, ok)
	})

	t.Run("it should delete a key correctly even with a collision", func(t *testing.T) {
		hashGeneratorFunc := func(key string) uint64 {
			return 1
		}

		hashTable := hashtable.New(
			5,
			hashtable.WithCustomHashGeneratorFunc[string, string](hashGeneratorFunc),
		)

		hashTable.Set("name", "Daniel")
		name, ok := hashTable.Get("name")
		assert.Equal(t, "Daniel", name)
		assert.True(t, ok)

		hashTable.Set("whatever", "whatever")
		whatever, ok := hashTable.Get("whatever")
		assert.Equal(t, "whatever", whatever)
		assert.True(t, ok)

		hashTable.Delete("name")
		name, ok = hashTable.Get("name")
		assert.Equal(t, "", name)
		assert.False(t, ok)
	})

	t.Run("it should delete a key correctly even with many collisions", func(t *testing.T) {
		hashGeneratorFunc := func(key string) uint64 {
			return 1
		}

		hashTable := hashtable.New(
			5,
			hashtable.WithCustomHashGeneratorFunc[string, string](hashGeneratorFunc),
		)

		hashTable.Set("name", "Daniel")
		name, ok := hashTable.Get("name")
		assert.Equal(t, "Daniel", name)
		assert.True(t, ok)

		hashTable.Set("whatever", "whatever")
		whatever, ok := hashTable.Get("whatever")
		assert.Equal(t, "whatever", whatever)
		assert.True(t, ok)

		hashTable.Set("whatever2", "whatever2")
		whatever2, ok := hashTable.Get("whatever2")
		assert.Equal(t, "whatever2", whatever2)
		assert.True(t, ok)

		hashTable.Set("whatever3", "whatever3")
		whatever3, ok := hashTable.Get("whatever3")
		assert.Equal(t, "whatever3", whatever3)
		assert.True(t, ok)

		hashTable.Delete("whatever2")
		whatever2, ok = hashTable.Get("whatever2")
		assert.Equal(t, "", whatever2)
		assert.False(t, ok)

		whatever3, ok = hashTable.Get("whatever3")
		assert.Equal(t, "whatever3", whatever3)
		assert.True(t, ok)

		hashTable.Delete("thiskeydoesntexist")

		whatever3, ok = hashTable.Get("whatever3")
		assert.Equal(t, "whatever3", whatever3)
		assert.True(t, ok)
	})
}

func TestHashTable_ConcurrentOperations(t *testing.T) {
	hashTable := hashtable.New[int, int](50)

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			key := rand.Intn(50)
			value, _ := hashTable.Get(key)
			value++
			hashTable.Set(key, value)
		}()
	}

	wg.Wait()
}
