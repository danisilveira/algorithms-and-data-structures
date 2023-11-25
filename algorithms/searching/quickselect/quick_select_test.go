package quickselect_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/searching/quickselect"
	"github.com/stretchr/testify/assert"
)

func TestQuickSelect(t *testing.T) {
	tt := []struct {
		items          []int
		kthLowestValue int
		expected       int
	}{
		{
			items:          []int{5, 4, 3, 2, 1},
			kthLowestValue: 2,
			expected:       2,
		},

		{
			items:          []int{10, 1, 5, 8, 3, -4, 5, 8, 3, 4, 5, 7, -3},
			kthLowestValue: 1,
			expected:       -4,
		},

		{
			items:          []int{1, 2, 3, 4, 5},
			kthLowestValue: 5,
			expected:       5,
		},
	}

	for _, tc := range tt {
		value := quickselect.Select(tc.items, tc.kthLowestValue)
		assert.Equal(t, tc.expected, value)
	}
}
