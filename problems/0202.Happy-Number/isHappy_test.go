package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {

	var testcase = []struct {
		n        int
		expected bool
	}{
		{
			n:        19,
			expected: true,
		},
		{
			n:        202,
			expected: false,
		},
		{},
		{},
		{},
	}

	t.Run("0202.Happy Number", func(t *testing.T) {
		for _, test := range testcase {
			got := isHappy(test.n)
			assert.Equal(t, test.expected, got)
		}
	})
}
