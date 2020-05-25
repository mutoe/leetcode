package happy_number

// https://leetcode.com/problems/happy-number

// level: 1
// time: O(1) 0ms 100%
// space: O(1) 2.1M

// leetcode submit region begin(Prohibit modification and deletion)
func sumOfSquares(n int) (ret int) {
	for n > 0 {
		remainder := n % 10
		ret += remainder * remainder
		n /= 10
	}
	return
}

func isHappy(n int) bool {
	numsMap := make(map[int]bool)

	for n != 1 {
		if numsMap[n] == true {
			return false
		}
		numsMap[n] = true
		n = sumOfSquares(n)
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
