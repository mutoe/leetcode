package make_the_string_great

// https://leetcode.com/problems/make-the-string-great

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.2M 53.85%

// leetcode submit region begin(Prohibit modification and deletion)
func isDifferentCase(a, b int32) bool {
	different := 'A' - 'a'
	if a+different == b || b+different == a {
		return true
	}
	return false
}

func makeGood(s string) string {
	stack := make([]int32, len(s))
	l := 0
	for _, char := range s {
		if l > 0 && isDifferentCase(stack[l-1], char) {
			l--
		} else {
			stack[l] = char
			l++
		}
	}
	return string(stack[:l])
}

// leetcode submit region end(Prohibit modification and deletion)
