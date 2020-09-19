package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {

	var testcase = []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2},
		},
		{
			nums1:    []int{4,9,5},
			nums2:    []int{9,4,9,8,4},
			expected: []int{9,4},
		},
	}

	t.Run("0349.Intersection of Two Arrays", func(t *testing.T) {
		for _, test := range testcase {
			got := intersection(test.nums1, test.nums2)
			assert.Equal(t, test.expected, got)
		}
	})
}
