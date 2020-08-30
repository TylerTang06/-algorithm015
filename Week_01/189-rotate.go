package week01

func rotate(nums []int, k int) {
	if nums == nil || len(nums) <= 1 || k == 0 || k == len(nums) {
		return
	}
	count := 0
	for stIndex := 0; count < len(nums); stIndex++ {
		temp, nxtIndx := nums[stIndex], (k+stIndex)%len(nums)
		for stIndex != nxtIndx {
			nums[nxtIndx], temp = temp, nums[nxtIndx]
			nxtIndx = (nxtIndx + k) % len(nums)
			count++
		}
		nums[nxtIndx] = temp
		count++
	}
}
