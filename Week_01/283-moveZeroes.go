package week01

func moveZeroes(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] == 0 {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
	}
}
