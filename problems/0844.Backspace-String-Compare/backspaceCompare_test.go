package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackspaceCompare(t *testing.T) {

	var testcase = []struct {
		S        string
		T        string
		expected bool
	}{
		{
			S:        "ab#c",
			T:        "ad#c",
			expected: true,
		},
		{
			S:        "ab##",
			T:        "c#d#",
			expected: true,
		},
		{
			S:        "a##c",
			T:        "#a#c",
			expected: true,
		},
		{
			S:        "a#c",
			T:        "b",
			expected: false,
		},
	}

	t.Run("0844.Backspace String Compare", func(t *testing.T) {
		for _, test := range testcase {
			got := backspaceCompare(test.S, test.T)
			assert.Equal(t, test.expected, got)
		}
	})

}
