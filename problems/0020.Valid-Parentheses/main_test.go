package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcase = []struct {
	s        string
	expected bool
}{
	{
		s:        "()",
		expected: true,
	},
	{
		s:        "()[]{}",
		expected: true,
	},
	{
		s:        "(]",
		expected: false,
	},
	{
		s:        "([)]",
		expected: false,
	},
	{
		s:        "{[]}",
		expected: true,
	},
}

func TestIsValid(t *testing.T) {
	t.Run("0020. Valid Parentheses", func(t *testing.T) {
		for _, test := range testcase {
			got := isValid(test.s)
			assert.Equal(t, test.expected, got)
		}
	})
}
