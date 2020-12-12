package sum_of_absolute_differences_in_a_sorted_array

// https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array

// level: 2

// TODO: Time Limit Exceeded

// leetcode submit region begin(Prohibit modification and deletion)
func getSumAbsoluteDifferences(nums []int) []int {
	ret := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			diff := nums[i] - nums[j]
			if diff >= 0 {
				ret[i] += diff
				ret[j] += diff
			} else {
				ret[i] -= diff
				ret[j] -= diff
			}
		}
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
