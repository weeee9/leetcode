package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {

	var testcase = []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			nums:     []int{2, 7, 9, 3, 1},
			expected: 22,
		},
		{
			nums:     []int{2},
			expected: 2,
		},
		{
			nums:     []int{-1, -2},
			expected: -1,
		},
	}

	t.Run("0053.Maximum SubArray", func(t *testing.T) {
		for _, test := range testcase {
			got := maxSubArray(test.nums)
			assert.Equal(t, test.expected, got)
		}
	})
}
