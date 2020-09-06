package week02

import (
	"container/list"

	"github.com/TylerTang06/-algorithm015/util"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *util.Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	myQue := list.New()
	myQue.PushBack(root)
	for myQue.Len() > 0 {
		len := myQue.Len()
		res := []int{}
		for len > 0 {
			node := myQue.Front().Value.(*util.Node)
			res = append(res, node.Val)

			for i := range node.Children {
				myQue.PushBack(node.Children[i])
			}
			myQue.Remove(myQue.Front())
			len--
		}
		result = append(result, res)
	}

	return result
}
