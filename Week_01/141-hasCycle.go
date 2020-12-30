package week01

import "github.com/TylerTang06/-algorithm015/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *util.ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	cur, nxt := head.Next, head.Next.Next
	for nxt.Next != nil && nxt.Next.Next != nil {
		if cur == nxt {
			return true
		}
		cur, nxt = cur.Next, nxt.Next.Next
	}

	return false
}
