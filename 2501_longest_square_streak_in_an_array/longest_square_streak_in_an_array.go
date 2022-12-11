package longest_square_streak_in_an_array

import "sort"

// https://leetcode.com/problems/longest-square-streak-in-an-array

// level: 2
// time: O(n) 487ms 100%
// space: O(n) 15.7M 100%

// leetcode submit region begin(Prohibit modification and deletion)
func longestSquareStreak(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] = 1
	}
	nums = []int{}
	result := -1
	for num := range m {
		nums = append(nums, num)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if m[num*num] > 0 {
			m[num] += m[num*num]
			if m[num] >= 2 && m[num] > result {
				result = m[num]
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
