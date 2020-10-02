package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestWord(t *testing.T) {
	var testcase = []struct {
		words    []string
		expected string
	}{
		{
			words:    []string{"w", "wo", "wor", "worl", "world"},
			expected: "world",
		},
		{
			words:    []string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
			expected: "apple",
		},
	}

	t.Run("0720.Longest Word in Dictionary", func(t *testing.T) {
		for _, test := range testcase {
			got := longestWord(test.words)
			assert.Equal(t, test.expected, got)
		}
	})
}
