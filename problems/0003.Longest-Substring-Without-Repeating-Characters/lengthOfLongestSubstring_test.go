package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var testcase = []struct {
		s        string
		expected int
	}{
		{
			s:        "abcabcbb",
			expected: 3,
		},
		{
			s:        "bbbbb",
			expected: 1,
		},
		{
			s:        "pwwkew",
			expected: 3,
		},
		{
			s:        "",
			expected: 0,
		},
	}

	t.Run("0003.Longest Substring Without Repeating Characters", func(t *testing.T) {
		for _, test := range testcase {
			got := lengthOfLongestSubstring(test.s)
			assert.Equal(t, test.expected, got)
		}
	})
}
