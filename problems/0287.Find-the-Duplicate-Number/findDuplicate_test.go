package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {

	var testcase = []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			nums:     []int{1, 1, 2},
			expected: 1,
		},
	}

	t.Run("0287.Find the Duplicate Number", func(t *testing.T) {
		for _, test := range testcase {
			got := findDuplicate(test.nums)
			assert.Equal(t, test.expected, got)
		}
	})

	t.Run("0287.Find the Duplicate Number: Sort", func(t *testing.T) {
		for _, test := range testcase {
			got := findDuplicateSort(test.nums)
			assert.Equal(t, test.expected, got)
		}
	})
}
