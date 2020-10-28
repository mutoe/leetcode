package excel_sheet_column_title

// https://leetcode.com/problems/excel-sheet-column-title

// level: 1
// time: O(log(n)) 0ms 100%
// space: O(1) 2M

// leetcode submit region begin(Prohibit modification and deletion)
func convertToTitle(n int) (ret string) {
	for n > 0 {
		char := n % 26
		n /= 26
		if char == 0 {
			char = 26
			n--
		}
		ret = string(rune('A'-1+char)) + ret
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
