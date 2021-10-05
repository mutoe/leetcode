package minimum_moves_to_convert_string

// https://leetcode.com/problems/minimum-moves-to-convert-string

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.3M 19.44%

//leetcode submit region begin(Prohibit modification and deletion)
func minimumMoves(s string) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			ret++
			i += 2
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
