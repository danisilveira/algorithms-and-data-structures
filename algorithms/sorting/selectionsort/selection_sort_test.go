package selectionsort_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/sorting/selectionsort"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort_ShouldBeAbleToSortAnArrayOfNumbers(t *testing.T) {
	items := []int{5, 4, 3, 2, 1}
	selectionsort.Sort(items)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, items)
}

func TestSelectionSort_ShouldBeAbleToSortAnArrayAlreadySorted(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	selectionsort.Sort(items)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, items)
}

func TestSelectionSort_ShouldBeAbleToSortAnArrayOfFloats(t *testing.T) {
	items := []float64{5.5, 4.4, 3.3, 2.2, 1.1}
	selectionsort.Sort(items)
	assert.Equal(t, []float64{1.1, 2.2, 3.3, 4.4, 5.5}, items)
}

func TestSelectionSort_ShouldBeAbleToSortAnArrayOfStrings(t *testing.T) {
	items := []string{"e", "d", "c", "b", "a"}
	selectionsort.Sort(items)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, items)
}
