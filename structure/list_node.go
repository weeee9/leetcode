package structure

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// IntsToList convert slice of int
// to List
func IntsToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	l := &ListNode{}
	tmp := l
	for _, n := range nums {
		tmp.Next = &ListNode{Val: n}
		tmp = tmp.Next
	}
	return l.Next
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
