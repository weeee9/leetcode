package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weeee9/leetcode/structure"
)

func TestAddTwoNumbers(t *testing.T) {

	var testcases = []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{
			l1:       []int{},
			l2:       []int{},
			expected: []int{},
		},
		{
			l1:       []int{1},
			l2:       []int{1},
			expected: []int{2},
		},
		{
			l1:       []int{1, 2, 3, 4},
			l2:       []int{1, 2, 3, 4},
			expected: []int{2, 4, 6, 8},
		},
		{
			l1:       []int{1, 2, 3, 4, 5},
			l2:       []int{1, 2, 3, 4, 5},
			expected: []int{2, 4, 6, 8, 0, 1},
		},
		{
			l1:       []int{9, 9, 9, 9, 9},
			l2:       []int{1},
			expected: []int{0, 0, 0, 0, 0, 1},
		},
		{
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			l1:       []int{1, 8, 3},
			l2:       []int{7, 1},
			expected: []int{8, 9, 3},
		},
	}

	t.Run("0002. Add Two Numbers", func(t *testing.T) {
		for _, test := range testcases {
			l1 := structure.IntsToList(test.l1)
			l2 := structure.IntsToList(test.l2)
			got := AddTwoNumbers(l1, l2)
			assert.Equal(t, test.expected, structure.List2Ints(got))
		}
	})
}
