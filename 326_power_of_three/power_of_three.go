package power_of_three

// https://leetcode.com/problems/power-of-three

// time: O(log(n)) 32ms 35.81%
// space: O(1) 6.1M

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	powerOfThree := 1
	for powerOfThree <= n {
		if powerOfThree == n {
			return true
		}
		powerOfThree *= 3
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)
