package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {

	var testcase = []struct {
		height   []int
		expected int
	}{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
	}

	t.Run("0011.Container With Most Water", func(t *testing.T) {
		for _, test := range testcase {
			got := maxArea(test.height)
			assert.Equal(t, test.expected, got)
		}
	})
}
