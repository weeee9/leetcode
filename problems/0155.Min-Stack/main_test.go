package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	t.Run("0155.Min Stack", func(t *testing.T) {
		stack := Constructor()
		stack.Push(-2)
		fmt.Printf("stack: %v\n", stack)
		stack.Push(0)
		fmt.Printf("stack: %v\n", stack)
		stack.Push(-3)
		fmt.Printf("stack: %v\n", stack)
		assert.Equal(t, -3, stack.GetMin())
		stack.Pop()
		fmt.Printf("stack: %v\n", stack)

		assert.Equal(t, 0, stack.Top())

		assert.Equal(t, -2, stack.GetMin())
	})

	t.Run("0155.Min Stack", func(t *testing.T) {
		stack := Constructor()
		stack.Push(2)
		fmt.Printf("stack: %v\n", stack)
		stack.Push(0)
		fmt.Printf("stack: %v\n", stack)
		stack.Push(3)
		fmt.Printf("stack: %v\n", stack)
		stack.Push(0)
		fmt.Printf("stack: %v\n", stack)
		assert.Equal(t, 0, stack.GetMin())
		stack.Pop()
		fmt.Printf("stack: %v\n", stack)
		fmt.Printf("stack: %v\n", stack)
		assert.Equal(t, 0, stack.GetMin())
		stack.Pop()
		fmt.Printf("stack: %v\n", stack)
		assert.Equal(t, 0, stack.GetMin())
		stack.Pop()
		fmt.Printf("stack: %v\n", stack)
		assert.Equal(t, 2, stack.GetMin())
	})
}
