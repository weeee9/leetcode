package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUgly(t *testing.T) {

	var testcase = []struct {
		num      int
		expected bool
	}{
		{
			num:      6,
			expected: true,
		},
		{
			num:      8,
			expected: true,
		},
		{
			num:      14,
			expected: false,
		},
	}

	t.Run("0263.Ugly Number", func(t *testing.T) {
		for _, test := range testcase {
			got := isUgly(test.num)
			assert.Equal(t, test.expected, got)
		}
	})
}
