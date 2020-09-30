package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixReshape(t *testing.T) {

	var testcase = []struct {
		nums     [][]int
		r        int
		c        int
		expected [][]int
	}{
		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 1,
			c: 4,
			expected: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 2,
			c: 4,
			expected: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}

	t.Run("0566.Reshape the Matrix", func(t *testing.T) {
		for _, test := range testcase {
			got := matrixReshape(test.nums, test.r, test.c)
			assert.Equal(t, test.expected, got)
		}
	})
}
