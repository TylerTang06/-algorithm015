package week01

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) <= 1 {
		return len(nums)
	}
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i] {
			index++
			if index != i {
				nums[index] = nums[i]
			}
		}
	}
	return index + 1
}
