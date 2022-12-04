package pascals_triangle_ii

// https://leetcode.com/problems/pascals-triangle-ii

// level: 1
// time: O(n^2) 3ms 21.07%
// space: O(n^2) 2.1M 39.89%

// leetcode submit region begin(Prohibit modification and deletion)
func getRow(rowIndex int) []int {
	dp := make([][]int, rowIndex+1)
	dp[0] = []int{1}
	for i := 1; i <= rowIndex; i++ {
		dp[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp[rowIndex]
}

//leetcode submit region end(Prohibit modification and deletion)
