package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplifyPath(t *testing.T) {
	var testcase = []struct {
		path     string
		expected string
	}{
		{
			path:     "/home/",
			expected: "/home",
		},
		{
			path:     "/../",
			expected: "/",
		},
		{
			path:     "/home//foo/",
			expected: "/home/foo",
		},
		{
			path:     "/a/./b/../../c/",
			expected: "/c",
		},
		{
			path:     "/a/../../b/../c//.//",
			expected: "/c",
		},
		{
			path:     "/a//b////c/d//././/..",
			expected: "/a/b/c",
		},
	}

	t.Run("0071.Simplify Path", func(t *testing.T) {
		for _, test := range testcase {
			got := simplifyPath(test.path)
			assert.Equal(t, test.expected, got)
		}
	})
}
