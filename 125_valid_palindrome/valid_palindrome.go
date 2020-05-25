package valid_palindrome

import (
	"strings"
)

// https://leetcode.com/problems/valid-palindrome

// // level: 1
// time: O(n*log(n)) 0ms 100%
// space: O(1) 2.9M

// a-z: 97-122
// A-Z: 65-90
// 0-9: 48-57

// leetcode submit region begin(Prohibit modification and deletion)
func isSupportedChars(char byte) bool {
	if char >= '0' && char <= '9' || char >= 'a' && char <= 'z' {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !isSupportedChars(s[i]) {
			i++
			continue
		}
		if !isSupportedChars(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)
