package leetcode

import (
	"github.com/weeee9/leetcode/structure"
)

// ListNode ...
type ListNode = structure.ListNode

// AddTwoNumbers ...
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}

	n1, n2 := 0, 0
	carry, currnet := 0, head

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		currnet.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		currnet = currnet.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
