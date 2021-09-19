package count_number_of_pairs_with_absolute_difference_k

// https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k

// level: 1
// time: O(n^2) 4ms 100%
// space: O(1) 3.2M 100%

//leetcode submit region begin(Prohibit modification and deletion)
func countKDifference(nums []int, k int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			diff := nums[j] - nums[i]
			if diff == k || -diff == k {
				ret++
			}
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
