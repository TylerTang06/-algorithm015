package week02

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{}
	}

	myMap := map[int]int{}
	for index, value := range nums {
		if _, ok := myMap[target-value]; ok {
			return []int{myMap[target-value], index}
		}
		myMap[value] = index
	}

	return []int{}
}
