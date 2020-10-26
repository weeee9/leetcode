package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	var testcase = []struct {
		A        [][]int
		expected [][]int
	}{
		{
			A: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
		{
			A: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}

	t.Run("0867. Transpose Matrix", func(t *testing.T) {
		for _, test := range testcase {
			got := transpose(test.A)
			assert.Equal(t, test.expected, got)
		}
	})
}
