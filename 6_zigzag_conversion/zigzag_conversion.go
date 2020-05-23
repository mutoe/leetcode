package zigzag_conversion

// https://leetcode.com/problems/zigzag-conversion

// time: O(n) 12ms 51.98%
// space: O(n) 6.4M

// leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) (ret string) {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	for i, sign := 0, 1; len(s) > 0; i += sign {
		rows[i] += string(s[0])
		s = s[1:]
		if sign > 0 && i >= numRows-1 {
			sign = -1
		} else if sign < 0 && i <= 0 {
			sign = 1
		}
	}
	for _, row := range rows {
		ret += row
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
