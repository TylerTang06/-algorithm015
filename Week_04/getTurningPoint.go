package week04

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
		if nums[mid] <= nums[l] && nums[mid] <= nums[r] {
			return mid
		}
		if nums[mid] > nums[l] {
			l = mid
		}
		if nums[mid] < nums[r] {
			r = mid
		}
	}
	return -1
}
