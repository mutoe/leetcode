package arranging_coins

// https://leetcode.com/problems/arranging-coins

// level: 1
// time: O(log(n)) 8ms 38.46%
// space: O(1) 2.2M

// leetcode submit region begin(Prohibit modification and deletion)
func arrangeCoins(n int) (ret int) {
	ret = 1
	for i := 0; i < n; i++ {
		n -= i
		if n-i < 0 {
			return i
		}
	}
	return n
}

// leetcode submit region end(Prohibit modification and deletion)
