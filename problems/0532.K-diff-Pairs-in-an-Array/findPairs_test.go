package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPairs(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{3, 1, 4, 1, 5},
			k:        2,
			expected: 2,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			k:        1,
			expected: 4,
		},
	}
	t.Run("0532.K-diff Pairs in an Array", func(t *testing.T) {
		for _, test := range testcases {
			got := findPairs(test.nums, test.k)
			assert.Equal(t, test.expected, got)
		}
	})
}
