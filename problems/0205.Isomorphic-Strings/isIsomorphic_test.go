package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsomorphic(t *testing.T) {

	var testcase = []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "egg",
			t:        "add",
			expected: true,
		},
		{
			s:        "foo",
			t:        "bar",
			expected: false,
		},
		{
			s:        "paper",
			t:        "title",
			expected: true,
		},
	}

	for _, test := range testcase {
		got := isIsomorphic(test.s, test.t)
		assert.Equal(t, test.expected, got)
	}
}
