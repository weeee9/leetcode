package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	var testcase = []struct {
		haystack string
		needle   string
		expected int
	}{
		{
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
		{
			haystack: "aaaaa",
			needle:   "baa",
			expected: -1,
		},
		{
			haystack: "a",
			needle:   "a",
			expected: 0,
		},
		{
			haystack: "mississippi",
			needle:   "pi",
			expected: 9,
		},
	}

	t.Run("0028.Implement strStr()", func(t *testing.T) {
		for _, test := range testcase {
			got := strStr(test.haystack, test.needle)
			assert.Equal(t, test.expected, got)
		}
	})
}
