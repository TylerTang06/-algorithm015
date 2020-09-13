package week03

import "sort"

func permuteUnique(nums []int) [][]int {
	if nums == nil || len(nums) <= 1 {
		return [][]int{nums}
	}

	// sort
	sort.Ints(nums)
	used := make([]bool, len(nums))
	for i := 0; i < len(used); i++ {
		used[i] = false
	}

	return recursionPU(nums, []int{}, used, [][]int{})
}

func recursionPU(nums, path []int, used []bool, res [][]int) [][]int {
	if len(nums) == len(path) {
		res = append(res, path)
		return res
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		// !used[i-1]
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		if !used[i] {
			used[i] = true
			path = append(path, nums[i])
			res = recursionPU(nums, path, used, res)
			used[i] = false
			path = append([]int{}, path[:len(path)-1]...)
		}
	}
	return res
}
