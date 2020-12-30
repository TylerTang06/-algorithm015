package week01

import (
	"github.com/TylerTang06/-algorithm015/util"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	l := &util.ListNode{Next: nil}
	cur := l
	for l2 != nil {
		if l1 == nil {
			l1, l2 = l2, l1
			break
		}
		if l1.Val < l2.Val {
			cur.Next, l1 = l1, l1.Next
		} else {
			cur.Next, l2 = l2, l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	return l.Next
}
