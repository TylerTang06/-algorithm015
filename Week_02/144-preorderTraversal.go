package week02

import (
	"container/list"

	"github.com/TylerTang06/-algorithm015/util"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *util.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	myStack := list.New()
	myStack.PushBack(root)
	for myStack.Len() > 0 {
		node := myStack.Back().Value.(*util.TreeNode)
		myStack.Remove(myStack.Back())
		res = append(res, node.Val)
		if node.Right != nil {
			myStack.PushBack(node.Right)
		}
		if node.Left != nil {
			myStack.PushBack(node.Left)
		}
	}

	return res
}
