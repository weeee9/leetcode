package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestoreString(t *testing.T) {
	var testcase = []struct {
		s        string
		indices  []int
		expected string
	}{
		{
			s:        "codeleet",
			indices:  []int{4, 5, 6, 7, 0, 2, 1, 3},
			expected: "leetcode",
		},
		{
			s:        "abc",
			indices:  []int{0, 1, 2},
			expected: "abc",
		},
		{
			s:        "aiohn",
			indices:  []int{3, 1, 4, 2, 0},
			expected: "nihao",
		},
		{
			s:        "aaiougrt",
			indices:  []int{4, 0, 2, 6, 7, 3, 1, 5},
			expected: "arigatou",
		},
		{
			s:        "art",
			indices:  []int{1, 0, 2},
			expected: "rat",
		},
	}

	t.Run("1528.Shuffle String", func(t *testing.T) {
		for _, test := range testcase {
			got := restoreString(test.s, test.indices)
			assert.Equal(t, test.expected, got)
		}
	})
}
