package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("0255.Implement Stack using Queues - 1", func(t *testing.T) {
		stack := Constructor()
		assert.Equal(t, true, stack.Empty())
		stack.Push(2)
		stack.Push(10)
		assert.Equal(t, 10, stack.Pop())
		assert.Equal(t, 2, stack.Top())
		assert.Equal(t, false, stack.Empty())
	})

	t.Run("0255.Implement Stack using Queues - 2", func(t *testing.T) {
		stack := Constructor()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		assert.Equal(t, 3, stack.Top())
	})

}
