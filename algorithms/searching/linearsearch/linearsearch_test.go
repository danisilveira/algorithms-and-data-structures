package linearsearch_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/searching/linearsearch"
	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	t.Run("it should return the correct index", func(t *testing.T) {
		index := linearsearch.Search([]int{1, 2, 3, 4, 5}, 5)
		assert.Equal(t, 4, index)
	})

	t.Run("it should return -1 if the element isn't on the list", func(t *testing.T) {
		index := linearsearch.Search([]int{1, 2, 3, 4, 5}, 6)
		assert.Equal(t, -1, index)
	})
}
