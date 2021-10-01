package maximum_difference_between_increasing_elements

// https://leetcode.com/problems/maximum-difference-between-increasing-elements

// level: 1
// time: O(n^2) 4ms 41.43%
// space: O(1) 2.6M 37.14%

//leetcode submit region begin(Prohibit modification and deletion)
func maximumDifference(nums []int) int {
	n := len(nums)
	max := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				diff := nums[j] - nums[i]
				if diff > max {
					max = diff
				}
			}
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
