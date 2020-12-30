package week01

import "github.com/TylerTang06/-algorithm015/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	cur, nxt := head.Next, head.Next.Next
	for nxt.Next != nil && nxt.Next.Next != nil {
		if cur == nxt {
			nxt = head
			break
		}
		cur, nxt = cur.Next, nxt.Next.Next
	}

	for nxt != nil && nxt != cur {
		nxt, cur = nxt.Next, cur.Next
	}

	return nxt
}
