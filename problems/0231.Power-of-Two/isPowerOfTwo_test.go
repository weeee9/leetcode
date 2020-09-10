package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {

	var testcase = []struct {
		n        int
		expected bool
	}{
		{
			n:        1,
			expected: true,
		},
		{
			n:        16,
			expected: true,
		},
		{
			n:        218,
			expected: false,
		},
	}

	t.Run("0231.Power of Two", func(t *testing.T) {
		for _, test := range testcase {
			got := isPowerOfTwo(test.n)
			assert.Equal(t, test.expected, got)
		}
	})
}
