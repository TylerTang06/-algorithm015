package week01

func moveZeroes(nums []int) {
	for nums == nil || len(nums) <= 1 {
		return
	}

	zeroIndex, valIndex := 0, 0
	for valIndex < len(nums) {
		if nums[valIndex] != 0 {
			nums[zeroIndex] = nums[valIndex]
			zeroIndex++
		}
		valIndex++
	}
	for zeroIndex < len(nums) {
		nums[zeroIndex] = 0
		zeroIndex++
	}
}
