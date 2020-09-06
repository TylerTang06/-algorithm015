package week02

import (
	"container/heap"
)

var myMap = map[int]int{}

type MyMapHeap []int

func (a MyMapHeap) Len() int {
	return len(a)
}
func (a MyMapHeap) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a MyMapHeap) Less(i, j int) bool {
	return myMap[a[i]] > myMap[a[j]]
}

func (a *MyMapHeap) Push(value interface{}) {
	*a = append(*a, value.(int))
}

func (a *MyMapHeap) Pop() interface{} {
	old := *a
	value := old[0]
	*a = old[1:]

	return value
}

func topKFrequent(nums []int, k int) []int {
	if nums == nil && len(nums) < k {
		return []int{}
	}
	for _, val := range nums {
		if _, ok := myMap[val]; ok {
			myMap[val]++
		} else {
			myMap[val] = 1
		}
	}
	myHeap := &MyMapHeap{}
	heap.Init(myHeap)
	for key, val := range myMap {
		if len(*myHeap) < k {
			heap.Push(myHeap, key)
		} else {
			if val > myMap[(*myHeap)[k-1]] {
				heap.Pop(myHeap)
				heap.Push(myHeap, key)
			}
		}
	}
	return *myHeap
}
