package mergesort_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/sorting/mergesort"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	tt := []struct {
		items    []int
		expected []int
	}{
		{
			items:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},

		{
			items:    []int{10, 1, 5, 8, 3, -4, 5, 8, 3, 4, 5, 7, -3},
			expected: []int{-4, -3, 1, 3, 3, 4, 5, 5, 5, 7, 8, 8, 10},
		},

		{
			items:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tt {
		value := mergesort.Sort(tc.items)
		assert.Equal(t, tc.expected, value)
	}
}
