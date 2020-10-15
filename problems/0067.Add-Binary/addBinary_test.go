package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	var testcase = []struct {
		a        string
		b        string
		expected string
	}{
		{
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
	}

	t.Run("0067.Add Binary", func(t *testing.T) {
		for _, test := range testcase {
			got := addBinary(test.a, test.b)
			assert.Equal(t, test.expected, got)
		}
	})
}
