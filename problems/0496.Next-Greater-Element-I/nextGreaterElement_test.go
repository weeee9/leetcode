package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {

	var testcase = []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			nums1:    []int{4, 1, 2},
			nums2:    []int{1, 3, 4, 2},
			expected: []int{-1, 3, -1},
		},
		{
			nums1:    []int{2, 4},
			nums2:    []int{1, 2, 3, 4},
			expected: []int{3, -1},
		},
	}

	t.Run("0496.Next Greater Number", func(t *testing.T) {
		for _, test := range testcase {
			got := nextGreaterElement(test.nums1, test.nums2)
			assert.Equal(t, test.expected, got)
		}
	})
}
