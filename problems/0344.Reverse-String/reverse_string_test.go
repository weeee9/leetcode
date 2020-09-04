package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcase = []struct {
	s        []byte
	expected []byte
}{
	{
		s:        []byte{'h', 'e', 'l', 'l', 'o'},
		expected: []byte{'o', 'l', 'l', 'e', 'h'},
	},
	{
		s:        []byte{'H', 'a', 'n', 'n', 'a', 'h'},
		expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
	},
}

func TestReverseString(t *testing.T) {
	t.Run("0344.Reverse String", func(t *testing.T) {
		for _, test := range testcase {
			reverseString(test.s)
			assert.Equal(t, test.expected, test.s)
		}
	})
}
