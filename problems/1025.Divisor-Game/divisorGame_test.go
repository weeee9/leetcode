package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisorGame(t *testing.T) {
	var testcase = []struct {
		N        int
		expected bool
	}{
		{
			N:        2,
			expected: true,
		},
		{
			N:        3,
			expected: false,
		},
		{
			N:        10,
			expected: true,
		},
		{
			N:        9,
			expected: false,
		},
	}

	t.Run("1025.Divisor Game", func(t *testing.T) {
		for _, test := range testcase {
			got := divisorGame(test.N)
			assert.Equal(t, test.expected, got)
		}
	})
}
