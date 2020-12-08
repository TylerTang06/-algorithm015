package week01

import "sort"

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) <= 2 {
		return [][]int{}
	}

	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[j] + nums[k]
			if sum == -nums[i] {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > -nums[i] {
				k--
			} else {
				j++
			}
		}
	}

	return res
}
