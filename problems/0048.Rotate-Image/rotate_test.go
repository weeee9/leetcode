package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {

	var testcase = []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 9},
			},
		},
	}

	t.Run("0048.Rotate Image", func(t *testing.T) {
		for _, test := range testcase {
			rotate(test.matrix)
			assert.Equal(t, test.expected, test.matrix)
		}
	})
}
