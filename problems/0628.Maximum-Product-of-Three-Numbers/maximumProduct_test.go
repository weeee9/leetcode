package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumProduct(t *testing.T) {
	var testcase = []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 3},
			expected: 6,
		},
		{
			nums:     []int{1, 2, 3, 4},
			expected: 24,
		},
		{
			nums:     []int{3, -1, 4},
			expected: -12,
		},
		{
			nums:     []int{2, 3, -2, 4},
			expected: 24,
		},
		{
			nums:     []int{-2, 0, -1, 2, 3, 1, 10},
			expected: 60,
		},
		{
			nums:     []int{1000, 1000, 0},
			expected: 0,
		},
	}

	for _, test := range testcase {
		got := maximumProduct(test.nums)
		assert.Equal(t, test.expected, got)
	}
}
