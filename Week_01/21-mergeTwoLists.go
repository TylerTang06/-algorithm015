package week01

import (
	. "github.com/TylerTang06/-algorithm015/util"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l := &ListNode{Next: nil}
	head := l
	for l2 != nil {
		if l1 == nil {
			l1, l2 = l2, nil
			continue
		}
		if l1.Val < l2.Val {
			l.Next, l1 = l1, l1.Next
		} else {
			l.Next, l2 = l2, l2.Next
		}
		l = l.Next
	}
	if l1 != nil {
		l.Next = l1
	}
	return head.Next
}
