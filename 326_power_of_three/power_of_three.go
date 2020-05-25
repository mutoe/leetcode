package power_of_three

// https://leetcode.com/problems/power-of-three

// level: 1
// time: O(1) 8ms 99.32%
// space: O(1) 6.1M

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	const maxPowerOfThree = 4052555153018976267
	return maxPowerOfThree%n == 0
}

// leetcode submit region end(Prohibit modification and deletion)
