package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {

	var testcase = []struct {
		A        []int
		expected []int
	}{
		{
			A:        []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			A:        []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}

	t.Run("0977.Squares of a Sorted Array", func(t *testing.T) {
		for _, test := range testcase {
			got := sortedSquares(test.A)
			assert.Equal(t, test.expected, got)
		}
	})
}
