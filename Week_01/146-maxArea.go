package week01

func maxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}

	l, r := 0, len(height)-1
	res := 0
	for l < r {
		h := height[l]
		if height[l] > height[r] {
			h = height[r]
			r--
		} else {
			l++
		}

		if res < h*(r-l+1) {
			res = h * (r - l + 1)
		}
	}

	return res
}
