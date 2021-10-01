package maximum_difference_between_increasing_elements

// https://leetcode.com/problems/maximum-difference-between-increasing-elements

// level: 1
// time: O(n) 8ms 18.57%
// space: O(1) 2.8M 28.57%

//leetcode submit region begin(Prohibit modification and deletion)
func maximumDifference(nums []int) int {
	n := len(nums)
	max := -1
	i, j := 0, 1
	for i < n && j < n {
		if nums[j] < nums[i] {
			i, j = j, j+1
			continue
		}
		if diff := nums[j] - nums[i]; diff > max {
			max = diff
		}
		j++
	}
	if max == 0 {
		return -1
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
