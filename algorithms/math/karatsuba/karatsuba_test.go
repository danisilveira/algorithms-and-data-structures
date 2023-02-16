package karatsuba_test

import (
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/math/karatsuba"
	"github.com/stretchr/testify/assert"
)

func TestKaratsubaMultiplication(t *testing.T) {
	tt := []struct {
		x        int64
		y        int64
		expected int64
	}{
		{
			x:        5678,
			y:        1234,
			expected: 7006652,
		},

		{
			x:        99,
			y:        99,
			expected: 9801,
		},

		{
			x:        123,
			y:        2,
			expected: 246,
		},

		{
			x:        10000000,
			y:        0,
			expected: 0,
		},

		{
			x:        21792189,
			y:        1,
			expected: 21792189,
		},
	}

	for _, tc := range tt {
		value := karatsuba.Multiply(tc.x, tc.y)
		assert.Equal(t, tc.expected, value)
	}
}
