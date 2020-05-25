package longest_substring_without_repeating_characters

// https://leetcode.com/problems/longest-substring-without-repeating-characters

// level: 2
// time: O(n^2) 584ms 5.03%
// space: O(n) 7M

// leetcode submit region begin(Prohibit modification and deletion)
func hasDuplicateChar(str string) bool {
	charMap := make(map[int32]bool)
	for _, char := range str {
		if charMap[char] {
			return true
		}
		charMap[char] = true
	}
	return false
}

func lengthOfLongestSubstring(s string) (maxLen int) {
	strLen := len(s)
	if strLen < 2 {
		return strLen
	}
	l, r := 0, 1
	for l < strLen && r <= strLen {
		curLen := r - l
		if hasDuplicateChar(s[l:r]) {
			l++
			continue
		}
		if curLen > maxLen {
			maxLen = curLen
		}
		r++
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
