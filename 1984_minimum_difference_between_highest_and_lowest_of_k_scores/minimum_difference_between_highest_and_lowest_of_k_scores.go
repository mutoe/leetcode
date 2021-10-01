package minimum_difference_between_highest_and_lowest_of_k_scores

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores

// level: 1
// time: O(n*log(n)) 15ms 85.85%
// space: O(1) 5.5M 16.04M

//leetcode submit region begin(Prohibit modification and deletion)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func minimumDifference(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		return abs(nums[1] - nums[0])
	}
	sort.Ints(nums)
	min := math.MaxInt32
	for i := k - 1; i < len(nums); i++ {
		if d := nums[i] - nums[i-k+1]; d < min {
			min = d
		}
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)
