package cache_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/cache"
	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	lru := cache.NewLRU[string, string](3)

	evicted := lru.Set("animals/dog", "Dog")
	assert.False(t, evicted)

	evicted = lru.Set("animals/cat", "Cat")
	assert.False(t, evicted)

	evicted = lru.Set("animals/bird", "Bird")
	assert.False(t, evicted)

	evicted = lru.Set("animals/salamander", "Salamander")
	assert.True(t, evicted)

	_, ok := lru.Get("animals/dog")
	assert.False(t, ok)

	value, ok := lru.Get("animals/cat")
	assert.True(t, ok)
	assert.Equal(t, value, "Cat")

	evicted = lru.Set("animals/shark", "Shark")
	assert.True(t, evicted)

	_, ok = lru.Get("animals/bird")
	assert.False(t, ok)

	evicted = lru.Set("animals/salamander", "Salamander")
	assert.False(t, evicted)

	evicted = lru.Set("animals/lion", "Lion")
	assert.True(t, evicted)

	_, ok = lru.Get("animals/cat")
	assert.False(t, ok)
}
