package matrix_cells_in_distance_order

// https://leetcode.com/problems/matrix-cells-in-distance-order

// level: 1
// time: O(n^2) 36ms 15.38%
// space: O(n^2) 7.1M

// leetcode submit region begin(Prohibit modification and deletion)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) (ret [][]int) {
	m := map[int][][]int{}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			distance := manhattanDistance(i, j, r0, c0)
			m[distance] = append(m[distance], []int{i, j})
		}
	}

	for i := 0; i < R+C; i++ {
		for _, ints := range m[i] {
			ret = append(ret, ints)
		}
	}

	return
}

// leetcode submit region end(Prohibit modification and deletion)
