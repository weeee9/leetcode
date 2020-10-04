package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximum69Number(t *testing.T) {
	var testcase = []struct {
		num      int
		expected int
	}{
		{
			num:      9669,
			expected: 9969,
		},
		{
			num:      9996,
			expected: 9999,
		},
		{
			num:      9999,
			expected: 9999,
		},
	}

	t.Run("1323.Maximum 69 Number", func(t *testing.T) {
		for _, test := range testcase {
			got := maximum69Number(test.num)
			assert.Equal(t, test.expected, got)
		}
	})
}
