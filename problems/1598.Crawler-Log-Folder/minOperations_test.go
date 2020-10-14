package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	var testcase = []struct {
		logs     []string
		expected int
	}{
		{
			logs:     []string{"d1/", "d2/", "../", "d21/", "./"},
			expected: 2,
		},
		{
			logs:     []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			expected: 3,
		},
		{
			logs:     []string{"d1/", "../", "../", "../"},
			expected: 0,
		},
	}

	t.Run("1598.Crawler Log Folder", func(t *testing.T) {
		for _, test := range testcase {
			got := minOperations(test.logs)
			assert.Equal(t, test.expected, got)
		}
	})
}
