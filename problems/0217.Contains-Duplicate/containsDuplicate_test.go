package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {

	var testcase = []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	t.Run("0217.Contains Duplicate", func(t *testing.T) {
		for _, test := range testcase {
			got := containsDuplicate(test.nums)
			assert.Equal(t, test.expected, got)
		}
	})
}
