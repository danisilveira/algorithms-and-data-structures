package inversions_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/searching/inversions"
	"github.com/stretchr/testify/assert"
)

func TestCountInversions(t *testing.T) {
	tt := []struct {
		items    []int
		expected int
	}{
		{
			items:    []int{1, 3, 5, 2, 4, 6},
			expected: 3,
		},

		{
			items:    []int{8, 4, 2, 1},
			expected: 6,
		},

		{
			items:    []int{},
			expected: 0,
		},

		{
			items:    []int{1},
			expected: 0,
		},

		{
			items:    []int{2, 1},
			expected: 1,
		},
	}

	for _, tc := range tt {
		value := inversions.Count(tc.items)
		assert.Equal(t, tc.expected, value)
	}
}
