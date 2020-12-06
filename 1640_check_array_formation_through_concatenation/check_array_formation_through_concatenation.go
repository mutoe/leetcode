package check_array_formation_through_concatenation

// https://leetcode.com/problems/check-array-formation-through-concatenation

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.8M 32.39%

// leetcode submit region begin(Prohibit modification and deletion)
func canFormArray(arr []int, pieces [][]int) bool {
	indexes := make(map[int]int)
	for i, num := range arr {
		indexes[num] = i
	}
	for _, piece := range pieces {
		for i, j := indexes[piece[0]], 0; j < len(piece); i, j = i+1, j+1 {
			if i >= len(arr) || piece[j] != arr[i] {
				return false
			}
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
