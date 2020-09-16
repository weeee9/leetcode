package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalPoints(t *testing.T) {

	var testcase = []struct {
		ops      []string
		expected int
	}{
		{
			ops:      []string{"5", "2", "C", "D", "+"},
			expected: 30,
		},
		{
			ops:      []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			expected: 27,
		},
	}

	t.Run("0682.Baseball Game", func(t *testing.T) {
		for _, test := range testcase {
			got := calPoints(test.ops)
			assert.Equal(t, test.expected, got)
		}
	})
}
