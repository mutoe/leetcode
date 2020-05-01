package power_of_three

import "math"

// https://leetcode.com/problems/power-of-three

// time: O(log(n)) 16ms 93.92%
// space: O(1) 6.1M

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	log3n := math.Log(float64(n)) / math.Log(3)
	if log3n-math.Floor(log3n) < 1e-10 {
		return true
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)
