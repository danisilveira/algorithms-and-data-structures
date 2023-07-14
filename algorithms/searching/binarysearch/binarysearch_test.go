package binarysearch_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/searching/binarysearch"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Run("it should return the correct index", func(t *testing.T) {
		index, steps := binarysearch.Search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4)
		assert.Equal(t, 3, index)
		assert.Equal(t, 4, steps)
	})

	t.Run("it should return the correct index when the value is on the middle of the list", func(t *testing.T) {
		index, steps := binarysearch.Search([]int{1, 2, 3, 4, 5}, 3)
		assert.Equal(t, 2, index)
		assert.Equal(t, 1, steps)
	})

	t.Run("it should return the correct index when the value is on the first half of the list", func(t *testing.T) {
		index, steps := binarysearch.Search([]int{1, 2, 3, 4, 5}, 2)
		assert.Equal(t, 1, index)
		assert.Equal(t, 2, steps)
	})

	t.Run("it should return the correct index when the value is on the second half of the list", func(t *testing.T) {
		index, steps := binarysearch.Search([]int{1, 2, 3, 4, 5}, 5)
		assert.Equal(t, 4, index)
		assert.Equal(t, 2, steps)
	})

	t.Run("it should return -1 if the value isn't on the list", func(t *testing.T) {
		index, steps := binarysearch.Search([]int{1, 2, 3, 4, 5}, 6)
		assert.Equal(t, -1, index)
		assert.Equal(t, 3, steps)
	})
}
