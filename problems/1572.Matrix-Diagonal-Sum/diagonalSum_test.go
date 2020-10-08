package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalSum(t *testing.T) {
	var testcase = []struct {
		mat      [][]int
		expected int
	}{
		{
			mat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 25,
		},
		{
			mat: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			expected: 8,
		},
		{
			mat: [][]int{
				{5},
			},
			expected: 5,
		},
	}

	t.Run("1572.Matrix Diagonal Sum", func(t *testing.T) {
		for _, test := range testcase {
			got := diagonalSum(test.mat)
			assert.Equal(t, test.expected, got)
		}
	})
}
