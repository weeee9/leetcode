package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {

	var testcase = []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 2},
			expected: 2,
		},
		{
			nums:     []int{0, 0, 1, 1, 1, 1, 2, 3, 4, 4},
			expected: 5,
		},
		{
			nums:     []int{0, 0, 0, 0, 0},
			expected: 1,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
	}

	t.Run("0026.Remov Duplicates from Sorted Array", func(t *testing.T) {
		for _, test := range testcase {
			got := removeDuplicates(test.nums)
			assert.Equal(t, test.expected, got)
		}
	})
}
