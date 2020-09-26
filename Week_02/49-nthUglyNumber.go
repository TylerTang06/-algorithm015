package week02

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0

	for i := 1; i < n; i++ {
		dp[i] = min(dp[a]*2, dp[b]*3, dp[c]*5)

		// should remove the case of dp[2] * 3 == dp[3] * 2
		if dp[i] >= dp[a]*2 {
			a++
		}
		if dp[i] >= dp[b]*3 {
			b++
		}
		if dp[i] >= dp[c]*5 {
			c++
		}
	}

	return dp[n-1]
}

func min(x, y, z int) int {
	if x > y {
		x = y
	}
	if x > z {
		return z
	}

	return x
}
