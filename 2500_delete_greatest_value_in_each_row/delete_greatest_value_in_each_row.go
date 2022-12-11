package delete_greatest_value_in_each_row

import "sort"

// https://leetcode.com/problems/delete-greatest-value-in-each-row

// level: 1
// time: O(n^2) 32ms 100%
// space: O(1) 4.8M 100%

// leetcode submit region begin(Prohibit modification and deletion)
func deleteGreatestValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}
	result := 0
	for i := 0; i < len(grid[0]); i++ {
		max := -1
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
		result += max
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
