package power_of_two

// https://leetcode.com/problems/power-of-two

// level: 1
// time: O(log(n)) 0ms 100%
// space: O(1) 2.2M

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfTwo(n int) bool {
	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n /= 2
	}
	return n == 1
}

// leetcode submit region end(Prohibit modification and deletion)
