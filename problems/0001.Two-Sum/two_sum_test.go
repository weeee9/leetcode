package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {

	var testcases = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 2, 4},
			target:   5,
			expected: []int{0, 1},
		},
		{
			nums:     []int{0, 8, 7, 3, 3, 4, 2},
			target:   11,
			expected: []int{1, 3},
		},
		{
			nums:     []int{0, 3},
			target:   5,
			expected: nil,
		},
		{
			nums:     []int{0, 1},
			target:   1,
			expected: []int{0, 1},
		},
	}

	t.Run("0001. Two Sum", func(t *testing.T) {
		for _, test := range testcases {
			got := TwoSum(test.nums, test.target)
			assert.Equal(t, test.expected, got)
		}
	})
}
