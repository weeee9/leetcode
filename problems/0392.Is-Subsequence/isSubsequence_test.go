package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	var testcase = []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
	}

	t.Run("0392. Is Subsequence", func(t *testing.T) {
		for _, test := range testcase {
			got := isSubsequence(test.s, test.t)
			assert.Equal(t, test.expected, got)
		}
	})
}
