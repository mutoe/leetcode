package remove_outermost_parentheses

import "bytes"

// https://leetcode.com/problems/remove-outermost-parentheses

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.5M 90.91%

// leetcode submit region begin(Prohibit modification and deletion)
func removeOuterParentheses(S string) string {
	level := 0
	var bf bytes.Buffer
	for _, char := range S {
		if char == '(' {
			level++
			if level > 1 {
				bf.WriteByte('(')
			}
		} else {
			level--
			if level != 0 {
				bf.WriteByte(')')
			}
		}
	}
	return bf.String()
}

// leetcode submit region end(Prohibit modification and deletion)
