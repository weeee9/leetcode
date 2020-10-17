package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidMountainArray(t *testing.T) {
	var testcase = []struct {
		A        []int
		expected bool
	}{
		{
			A:        []int{2, 1},
			expected: false,
		},
		{
			A:        []int{3, 5, 5},
			expected: false,
		},
		{
			A:        []int{0, 3, 2, 1},
			expected: true,
		},
	}

	t.Run("0941. Valid Mountain Array", func(t *testing.T) {
		for _, test := range testcase {
			got := validMountainArray(test.A)
			assert.Equal(t, test.expected, got)
		}
	})
}
