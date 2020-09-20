package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {

	var testcase = []struct {
		s        string
		expected bool
	}{
		{
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			s:        "race a car",
			expected: false,
		},
	}

	t.Run("0125.Valid Palindrome", func(t *testing.T) {
		for _, test := range testcase {
			got := isPalindrome(test.s)
			assert.Equal(t, test.expected, got)
		}
	})

}
