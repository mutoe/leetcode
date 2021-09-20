package minimum_value_to_get_positive_step_by_step_sum

import "math"

// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.1M 33.33%

//leetcode submit region begin(Prohibit modification and deletion)
func minStartValue(nums []int) int {
	N := len(nums)
	sum := make([]int, N+1)
	for i := 0; i < N; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	min := math.MaxInt32
	for i := 0; i <= N; i++ {
		if sum[i] < min {
			min = sum[i]
		}
	}
	return 1 - min
}

//leetcode submit region end(Prohibit modification and deletion)
