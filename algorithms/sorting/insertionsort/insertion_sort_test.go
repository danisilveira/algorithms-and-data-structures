package insertionsort_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/sorting/insertionsort"
	"github.com/stretchr/testify/assert"
)

func TestInsertionSort_ShouldBeAbleToSortAnArrayOfNumbers(t *testing.T) {
	items := []int{5, 4, 3, 2, 1}
	insertionsort.Sort(items)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, items)
}

func TestInsertionSort_ShouldBeAbleToSortAnArrayAlreadySorted(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	insertionsort.Sort(items)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, items)
}

func TestInsertionSort_ShouldBeAbleToSortAnArrayOfFloats(t *testing.T) {
	items := []float64{5.5, 4.4, 3.3, 2.2, 1.1}
	insertionsort.Sort(items)
	assert.Equal(t, []float64{1.1, 2.2, 3.3, 4.4, 5.5}, items)
}

func TestInsertionSort_ShouldBeAbleToSortAnArrayOfStrings(t *testing.T) {
	items := []string{"e", "d", "c", "b", "a"}
	insertionsort.Sort(items)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, items)
}
