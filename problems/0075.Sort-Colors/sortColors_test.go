package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortColors(t *testing.T) {
	var testcase = []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			nums:     []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			nums:     []int{0},
			expected: []int{0},
		},
		{
			nums:     []int{1},
			expected: []int{1},
		},
	}

	t.Run("0075.Sort Colors", func(t *testing.T) {
		for _, test := range testcase {
			sortColors(test.nums)
			assert.Equal(t, test.expected, test.nums)
		}
	})

}
