package ugly_number_ii

// https://leetcode.com/problems/ugly-number-ii

// level: 2

//leetcode submit region begin(Prohibit modification and deletion)

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}

	dp := map[int]bool{}
	for i := 1; i <= 6; i++ {
		dp[i] = true
	}
	i := 7
	for u := 6; u < n; i++ {
		if (i%2 == 0 && dp[i/2]) || (i%3 == 0 && dp[i/3]) || (i%5 == 0 && dp[i/5]) {
			dp[i] = true
			u++
		}
	}
	return i - 1
}

//leetcode submit region end(Prohibit modification and deletion)
