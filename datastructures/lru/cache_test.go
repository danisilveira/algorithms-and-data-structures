package lru_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/lru"
	"github.com/stretchr/testify/assert"
)

func Test_LRU(t *testing.T) {
	cache := lru.NewCache[string, string](3)

	evicted := cache.Set("animals/dog", "Dog")
	assert.False(t, evicted)

	evicted = cache.Set("animals/cat", "Cat")
	assert.False(t, evicted)

	evicted = cache.Set("animals/bird", "Bird")
	assert.False(t, evicted)

	evicted = cache.Set("animals/salamander", "Salamander")
	assert.True(t, evicted)

	_, ok := cache.Get("animals/dog")
	assert.False(t, ok)

	value, ok := cache.Get("animals/cat")
	assert.True(t, ok)
	assert.Equal(t, value, "Cat")

	evicted = cache.Set("animals/shark", "Shark")
	assert.True(t, evicted)

	_, ok = cache.Get("animals/bird")
	assert.False(t, ok)

	evicted = cache.Set("animals/salamander", "Salamander")
	assert.False(t, evicted)

	evicted = cache.Set("animals/lion", "Lion")
	assert.True(t, evicted)

	_, ok = cache.Get("animals/cat")
	assert.False(t, ok)
}
