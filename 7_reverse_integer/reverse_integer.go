package reverse_integer

import "math"

// https://leetcode.com/problems/reverse-integer

// level: 1
// time: O(log(n)) 0ms 100%
// space: O(1) 2.2M

// leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) (ret int) {
	isPositive := x >= 0
	if !isPositive {
		x = -x
	}
	for x > 0 {
		ret = ret*10 + x%10
		x /= 10
	}

	if !isPositive {
		ret = -ret
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
