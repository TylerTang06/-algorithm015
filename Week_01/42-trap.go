package week01

func trap(height []int) int {
	if height == nil || len(height) <= 2 {
		return 0
	}

	maxL, maxR, sum := 0, 0, 0
	left, right := 1, len(height)-2
	for i := 1; i < len(height)-1; i++ {
		if height[left-1] < height[right+1] {
			if maxL < height[left-1] {
				maxL = height[left-1]
			}
			if maxL > height[left] {
				sum += (maxL - height[left])
			}
			left++
		} else {
			if maxR < height[right+1] {
				maxR = height[right+1]
			}
			if maxR > height[right] {
				sum += (maxR - height[right])
			}
			right--
		}
	}

	return sum
}
