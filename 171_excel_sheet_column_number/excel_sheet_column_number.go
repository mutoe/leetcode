package excel_sheet_column_number

// https://leetcode.com/problems/excel-sheet-column-number

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.2M

// leetcode submit region begin(Prohibit modification and deletion)
func titleToNumber(s string) (ret int) {
	for _, char := range s {
		ret = ret*26 + int(char-'A') + 1
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
