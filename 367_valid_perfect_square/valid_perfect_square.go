package valid_perfect_square

// https://leetcode.com/problems/valid-perfect-square

// level: 1
// time: O(log(n)) 0ms 100%
// space: O(1) 2M 11.54%

// leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 1, num
	for l <= r {
		mid := l + (r-l)>>1
		sq := mid * mid
		if sq == num {
			return true
		}
		if sq > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)
