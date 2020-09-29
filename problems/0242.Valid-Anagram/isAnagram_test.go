package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {

	var testcase = []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			s:        "rat",
			t:        "car",
			expected: false,
		},
	}

	t.Run("0242.Valid Anagram", func(t *testing.T) {
		for _, test := range testcase {
			got := isAnagram(test.s, test.t)
			assert.Equal(t, test.expected, got)
		}
	})
}
