package sqrtx

// https://leetcode.com/problems/sqrtx

// level: 1
// time: O(log(n)) 8ms 19.58%
// space: O(1) 2.2M 100%

// TODO: Need to review

// leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l, r := 1, x
	for l < r {
		mid := l + (r-l)>>1
		midSquare := mid * mid
		if midSquare == x {
			return mid
		} else if midSquare > x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l - 1
}

// leetcode submit region end(Prohibit modification and deletion)
