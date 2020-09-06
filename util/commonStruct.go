package util

import "container/heap"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type KthLargest struct {
	myHeap *MinIntHeap
	k      int
}

func NewKthLargest(k int) *KthLargest {
	myHeap := &MinIntHeap{}
	heap.Init(myHeap)
	return &KthLargest{
		myHeap: myHeap,
		k:      k,
	}
}

func (k *KthLargest) Add(val int) int {
	if len(*k.myHeap) < k.k || val > (*k.myHeap)[0] {
		if len(*k.myHeap) == k.k {
			heap.Pop(k.myHeap)
		}
		heap.Push(k.myHeap, val)
	}
	return (*k.myHeap)[0]
}

func (k *KthLargest) All() []int {
	return *k.myHeap
}
