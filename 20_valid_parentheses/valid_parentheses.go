package valid_parentheses

// https://leetcode.com/problems/valid-parentheses

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2M 38.92%

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]int32, 0)
	charMap := map[int32]int32{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	for _, char := range s {
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if charMap[c] != char {
				return false
			}

		}
	}

	return len(stack) == 0
}

// leetcode submit region end(Prohibit modification and deletion)
