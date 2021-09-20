package find_pivot_index

// https://leetcode.com/problems/find-pivot-index

// level: 1
// time: O(n) 12ms 100%
// space: O(n) 6.2M 97.47%

//leetcode submit region begin(Prohibit modification and deletion)
func pivotIndex(nums []int) int {
	N := len(nums)
	sum := make([]int, N+1)
	for i := 0; i < N; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	for i := 0; i < N; i++ {
		if sum[i] == sum[N]-sum[i+1] {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
