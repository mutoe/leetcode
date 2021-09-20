package running_sum_of_1d_array

// https://leetcode.com/problems/running-sum-of-1d-array

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.8M 41.34%

//leetcode submit region begin(Prohibit modification and deletion)
func runningSum(nums []int) []int {
	N := len(nums)
	sum := make([]int, N+1)
	for i := 0; i < N; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	return sum[1:]
}

//leetcode submit region end(Prohibit modification and deletion)
