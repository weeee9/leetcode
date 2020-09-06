package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLongPressedName(t *testing.T) {
	var testcase = []struct {
		name     string
		typed    string
		expected bool
	}{
		{
			name:     "alex",
			typed:    "aaleex",
			expected: true,
		},
		{
			name:     "alex",
			typed:    "alexxr",
			expected: false,
		},
		{
			name:     "saeed",
			typed:    "ssaaedd",
			expected: false,
		},
		{
			name:     "leelee",
			typed:    "lleeelee",
			expected: true,
		},
		{
			name:     "laiden",
			typed:    "laiden",
			expected: true,
		},
		{
			name:     "kikcxmvzi",
			typed:    "kiikcxxmmvvzz",
			expected: false,
		},
	}

	t.Run("0925.Long Pressed Name", func(t *testing.T) {
		for _, test := range testcase {

			got := isLongPressedName(test.name, test.typed)
			assert.Equal(t, test.expected, got)
		}
	})
}
