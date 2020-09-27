package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIdenticalPairs(t *testing.T) {
	var testcase = []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 3, 1, 1, 3},
			expected: 4,
		},
		{
			nums:     []int{1,1,1,1},
			expected: 6,
		},
		{
			nums:     []int{1,2,3},
			expected: 0,
		},
	}

	t.Run("1512. Number of Good Pairs", func(t *testing.T) {
		for _, test := range testcase {
			got := numIdenticalPairs(test.nums)
			assert.Equal(t, test.expected, got)
		}
	})
}
