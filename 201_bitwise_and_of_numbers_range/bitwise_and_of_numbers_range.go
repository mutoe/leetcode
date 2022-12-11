package bitwise_and_of_numbers_range

// https://leetcode.com/problems/bitwise-and-of-numbers-range

// level: 2
// time: O(n) 822ms 6.25%
// space: O(1) 5.3M 18.75%

// leetcode submit region begin(Prohibit modification and deletion)
func rangeBitwiseAnd(left int, right int) int {
	result := left
	for i := left + 1; i <= right; i++ {
		result &= i
		if result == 0 {
			return 0
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
