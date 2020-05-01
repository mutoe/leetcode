package add_digits

// https://leetcode.com/problems/add-digits

// time: O(1) 0ms 100%
// space: O(1) 2.2M

// leetcode submit region begin(Prohibit modification and deletion)
func addDigits(num int) int {
	for num > 9 {
		ret := 0
		for num > 0 {
			ret += num % 10
			num /= 10
		}
		num = ret
	}
	return num
}

// leetcode submit region end(Prohibit modification and deletion)
