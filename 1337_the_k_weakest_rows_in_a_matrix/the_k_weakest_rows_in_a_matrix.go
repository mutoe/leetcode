package the_k_weakest_rows_in_a_matrix

// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix

// level: 1
// time: O(n*m) 12ms 82.89%
// space: O(n) 5M 71.05%

// leetcode submit region begin(Prohibit modification and deletion)
func kWeakestRows(mat [][]int, k int) []int {
	ret := make([]int, 0)

	for j := 0; j <= len(mat[0]); j++ {
		for i := 0; i < len(mat); i++ {
			if len(ret) == k {
				return ret
			}
			if includes(ret, i) {
				continue
			}
			if j == len(mat[0]) || mat[i][j] == 0 {
				ret = append(ret, i)
			}
		}
	}
	return ret
}

func includes(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)
