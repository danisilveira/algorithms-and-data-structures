package hashtable_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/hashtable"
	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	hashTable := hashtable.New[string, string](5)

	hashTable.Set("name", "Daniel")
	name, ok := hashTable.Get("name")
	assert.Equal(t, "Daniel", name)
	assert.True(t, ok)
}

func TestHashTable_ShouldBeAbleToDealWithCollisionsCorrectly(t *testing.T) {
	hashGeneratorFunc := func(key string) int64 {
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
}

func TestHashTable_ShouldOverrideValueWithTheSameKey(t *testing.T) {
	hashTable := hashtable.New[string, string](5)

	hashTable.Set("name", "Daniel")
	name, ok := hashTable.Get("name")
	assert.Equal(t, "Daniel", name)
	assert.True(t, ok)

	hashTable.Set("name", "Lucas")
	name, ok = hashTable.Get("name")
	assert.Equal(t, "Lucas", name)
	assert.True(t, ok)
}

func TestHashTable_ShouldOverrideValueWithTheSameKeyEvenWithACollision(t *testing.T) {
	hashGeneratorFunc := func(key string) int64 {
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
}

func TestHashTable_ShouldReturnTheDefaultValueWhenKeyIsNotFound(t *testing.T) {
	hashTable := hashtable.New[string, string](5)

	hashTable.Set("name", "Daniel")
	name, ok := hashTable.Get("name")
	assert.Equal(t, "Daniel", name)
	assert.True(t, ok)

	key, ok := hashTable.Get("otherKey")
	assert.Equal(t, "", key)
	assert.False(t, ok)
}
