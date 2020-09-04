package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	var testcase = []struct {
		s        string
		expected string
	}{
		{
			s:        "hello",
			expected: "holle",
		},
		{
			s:        "leetcode",
			expected: "leotcede",
		},
		{
			s:        "aA",
			expected: "Aa",
		},
	}
	t.Run("0345.Reverse Vowels of a String", func(t *testing.T) {
		for _, test := range testcase {
			got := reverseVowels(test.s)
			assert.Equal(t, test.expected, got)
		}
	})
}
