package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	var testcase = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			nums:     []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			nums:     []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
	}

	t.Run("0167.Two Sum II", func(t *testing.T) {
		for _, test := range testcase {
			got := twoSum(test.nums, test.target)
			assert.Equal(t, test.expected, got)
		}
	})
}
