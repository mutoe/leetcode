package power_of_two

// https://leetcode.com/problems/power-of-two

// time: O(n) 4ms 36.17%
// space: O(1) 2.2M

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfTwo(n int) bool {
	if n > 1 && n%2 == 1 {
		return false
	}
	base := 1
	for {
		if n == base {
			return true
		}
		base *= 2
		if base > n {
			return false
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
