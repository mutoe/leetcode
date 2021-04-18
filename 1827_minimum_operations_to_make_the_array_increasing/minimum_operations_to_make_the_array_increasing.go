package minimum_operations_to_make_the_array_increasing

// https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing

// level: 1
// time: O(n) 12ms 100%
// space: O(1) 5.8M 100%

//leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums []int) int {
	last := nums[0]
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > last {
			last = nums[i]
		} else {
			count += last - nums[i] + 1
			last++
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
