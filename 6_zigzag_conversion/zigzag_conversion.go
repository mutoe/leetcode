package zigzag_conversion

// https://leetcode.com/problems/zigzag-conversion

// time: O(n) 8ms 74.58%
// space: O(n) 5.8M

// leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]byte, numRows)
	for si, i, sign := 0, 0, 1; si < len(s); {
		rows[i] = append(rows[i], s[si])
		if sign > 0 && i >= numRows-1 {
			sign = -1
		} else if sign < 0 && i <= 0 {
			sign = 1
		}
		si++
		i += sign
	}
	result := make([]byte, 0, len(s))
	for _, row := range rows {
		result = append(result, row...)
	}
	return string(result)
}

// leetcode submit region end(Prohibit modification and deletion)
