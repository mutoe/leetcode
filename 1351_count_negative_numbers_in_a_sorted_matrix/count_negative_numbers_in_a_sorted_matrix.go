package count_negative_numbers_in_a_sorted_matrix

// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix

// level: 1
// time: O(log(n)) 8ms 99.26%
// space: O(1) 6M 97.78%

// leetcode submit region begin(Prohibit modification and deletion)
func countNegatives(grid [][]int) int {
	ret := 0
	m, n := len(grid), len(grid[0])
	r := n - 1
	for i := 0; i < m; i++ {
		if r < 0 {
			ret += n
			continue
		}
		for l, mid := 0, 0; l <= r; mid = (l + r) >> 1 {
			if grid[i][mid] >= 0 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		ret += n - r - 1
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
