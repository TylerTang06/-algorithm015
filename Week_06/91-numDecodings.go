package week06

import "strconv"

func numDecodings(s string) int {
	if s == "" || s[0] == '0' {
		return 0
	}

	// dp[i] = dp[i-1] if s[i] == '0'
	// dp[i] += dp[i-2]
	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			dp[i] = dp[i-1]
		}
		dp[i] = dp[i-1]
		num, _ := strconv.Atoi(string(s[i-1]) + string(s[i]))
		if num <= 26 && num >= 10 {
			if i == 2 {
				dp[i] = dp[i-1]
			} else {
				dp[i] += dp[i-2]
			}
		}
	}

	return dp[len(s)-1]
}
