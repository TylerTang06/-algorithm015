package week01

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		n2, n1 = n2+n1, n2
	}

	return n2
}
