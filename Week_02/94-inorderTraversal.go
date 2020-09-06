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
func inorderTraversal(root *util.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	myStack := list.New()
	myMap := map[*util.TreeNode]int{}

	myStack.PushBack(root)
	myMap[root] = 1
	for myStack.Len() > 0 {
		node := myStack.Back().Value.(*util.TreeNode)
		myStack.Remove(myStack.Back())
		if 1 == myMap[node] {
			if node.Right != nil {
				myStack.PushBack(node.Right)
				myMap[node.Right] = 1
			}
			myStack.PushBack(node)
			myMap[node] = 2
			if node.Left != nil {
				myStack.PushBack(node.Left)
				myMap[node.Left] = 1
			}
		} else {
			res = append(res, node.Val)
		}
	}

	return res
}
