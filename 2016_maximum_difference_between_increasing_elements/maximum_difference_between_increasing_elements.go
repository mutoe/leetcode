package maximum_difference_between_increasing_elements

// https://leetcode.com/problems/maximum-difference-between-increasing-elements

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.6M 37.14%

//leetcode submit region begin(Prohibit modification and deletion)
func maximumDifference(nums []int) int {
	diff := -1
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min, max = num, num
		} else if num > max {
			max = num
			if d := max - min; d > diff {
				diff = d
			}
		}
	}
	return diff
}

//leetcode submit region end(Prohibit modification and deletion)
