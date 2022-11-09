package algorithms_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort_ShouldBeAbleToSortAnArrayOfNumbers(t *testing.T) {
	items := []int{5, 4, 3, 2, 1}
	algorithms.BubbleSort(items)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, items)
}

func TestBubbleSort_ShouldBeAbleToSortAnArrayAlreadySorted(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	algorithms.BubbleSort(items)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, items)
}

func TestBubbleSort_ShouldBeAbleToSortAnArrayOfFloats(t *testing.T) {
	items := []float64{5.5, 4.4, 3.3, 2.2, 1.1}
	algorithms.BubbleSort(items)
	assert.Equal(t, []float64{1.1, 2.2, 3.3, 4.4, 5.5}, items)
}

func TestBubbleSort_ShouldBeAbleToSortAnArrayOfStrings(t *testing.T) {
	items := []string{"e", "d", "c", "b", "a"}
	algorithms.BubbleSort(items)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, items)
}
