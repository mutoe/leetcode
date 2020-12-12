package count_negative_numbers_in_a_sorted_matrix

// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix

// level: 1
// time: O(log(n)) 16ms 69.63%
// space: O(1) 6M 97.78%

// leetcode submit region begin(Prohibit modification and deletion)
func countNegatives(grid [][]int) int {
	ret := 0
	pivot := len(grid[0]) - 1
	for i := 0; i < len(grid); i++ {
		l, r := 0, pivot
		for l <= r {
			mid := (l + r) >> 1
			if grid[i][mid] >= 0 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		pivot = r
		ret += len(grid[0]) - r - 1
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
