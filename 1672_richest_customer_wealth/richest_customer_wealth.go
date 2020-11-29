package richest_customer_wealth

// https://leetcode.com/problems/richest-customer-wealth

// level: 1
// time: O(n^2) 0ms 100%
// space: O(n) 3.3M

// leetcode submit region begin(Prohibit modification and deletion)
func maximumWealth(accounts [][]int) int {
	sums := make([]int, len(accounts))

	max := 0
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			sums[i] += accounts[i][j]
		}
		if sums[i] > max {
			max = sums[i]
		}
	}
	return max
}

// leetcode submit region end(Prohibit modification and deletion)
