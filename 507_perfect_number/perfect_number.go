package perfect_number

import "math"

// https://leetcode.com/problems/perfect-number

// level: 1
// time: O(log(n)) 0ms 100%
// space: O(1) 2M

// leetcode submit region begin(Prohibit modification and deletion)
func checkPerfectNumber(num int) bool {
	if num < 6 {
		return false
	}
	sqrtNum := int(math.Sqrt(float64(num)))
	divisorSum := 1
	for i := 2; i <= sqrtNum; i++ {
		if num%i == 0 {
			divisorSum += i + num/i
		}
	}
	return divisorSum == num
}

// leetcode submit region end(Prohibit modification and deletion)
