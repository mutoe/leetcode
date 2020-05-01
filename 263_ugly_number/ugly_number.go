package ugly_number

// https://leetcode.com/problems/ugly-number

// time: O(log(n)) 0ms 100%
// space: O(1) 2.1M

// leetcode submit region begin(Prohibit modification and deletion)
func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for num > 1 {
		if num%2 == 0 {
			num /= 2
			continue
		}
		if num%3 == 0 {
			num /= 3
			continue
		}
		if num%5 == 0 {
			num /= 5
			continue
		}
		return false
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
