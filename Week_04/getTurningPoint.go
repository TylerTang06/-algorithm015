package week04

import "fmt"

func getTurningPoint(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	if nums[0] < nums[len(nums)-1] {
		return 0
	}

	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		fmt.Println(l, mid, r)
		if nums[mid] >= nums[l] && nums[mid] <= nums[r] {
			return nums[l]
		}
		if nums[mid] > nums[l] {
			l = mid + 1
		}
		if nums[mid] < nums[r] {
			r = mid - 1
		}
		fmt.Println(l, r)
	}

	return nums[l]
}
