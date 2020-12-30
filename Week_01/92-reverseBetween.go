package week01

import "github.com/TylerTang06/-algorithm015/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *util.ListNode, m int, n int) *util.ListNode {
	if m == n {
		return head
	}

	// notice: bef.Next = head
	curH, bef := head, &util.ListNode{Next: head}
	for i := 0; i < m-1 && curH != nil; i++ {
		curH, bef = curH.Next, bef.Next
	}
	bef.Next = nil // disconnet bef and curH

	cur, curB := curH.Next, curH // curB record the end of curH
	curH.Next = nil              // get the firt node of curH
	for i := m; i < n && cur != nil; i++ {
		nxt := cur.Next
		cur.Next, curH = curH, cur
		cur = nxt
	}

	if m == 1 {
		head, curB.Next = curH, cur
	} else {
		bef.Next, curB.Next = curH, cur
	}

	return head
}
