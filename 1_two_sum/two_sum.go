package two_sum

// https://leetcode.com/problems/two-sum

// level: 1
// time: O(n*log(n)) 36ms 30.28%
// space: O(1) 2.9M

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// leetcode submit region end(Prohibit modification and deletion)
