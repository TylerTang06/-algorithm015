package week01

func plusOne(digits []int) []int {
	if digits == nil || len(digits) == 0 {
		return digits
	}

	digits[len(digits)-1]++
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+carry < 10 {
			digits[i] += carry
			return digits
		}
		digits[i] = digits[i] + carry - 10
		carry = 1
	}
	if carry == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}
