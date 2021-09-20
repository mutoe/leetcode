package find_the_middle_index_in_array

// https://leetcode.com/problems/find-the-middle-index-in-array

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.5M 59.26%

//leetcode submit region begin(Prohibit modification and deletion)
func findMiddleIndex(nums []int) int {
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
