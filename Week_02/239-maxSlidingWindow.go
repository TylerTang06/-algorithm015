package week02

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) < k {
		return []int{}
	}

	res := []int{}
	// can use arry to be queue
	myQue := list.New()
	for i := 0; i < len(nums); i++ {
		if i > k-1 && myQue.Front().Value.(int) == i-k {
			myQue.Remove(myQue.Front())
		}
		// nums[i] compare with back value of Queue
		for myQue.Len() > 0 && nums[myQue.Back().Value.(int)] <= nums[i] {
			myQue.Remove(myQue.Back())
		}

		myQue.PushBack(i)
		if i >= k-1 {
			res = append(res, nums[myQue.Front().Value.(int)])
		}
	}

	return res
}
