package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {

	var testcase = []struct {
		nums     []int
		val      int
		expected int
	}{
		{
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
		{
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 6,
		},
	}

	t.Run("0027.Remove Element", func(t *testing.T) {
		for _, test := range testcase {
			got := removeElement(test.nums, test.val)
			assert.Equal(t, test.expected, got)
		}
	})

}
