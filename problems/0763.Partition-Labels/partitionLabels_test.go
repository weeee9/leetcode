package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionLabels(t *testing.T) {
	var testcase = []struct {
		s        string
		expected []int
	}{
		{
			s:        "ababcbacadefegdehijhklij",
			expected: []int{9, 7, 8},
		},
	}

	t.Run("0763.Partition Labels", func(t *testing.T) {
		for _, test := range testcase {
			got := partitionLablels(test.s)
			assert.Equal(t, test.expected, got)
		}
	})
}
